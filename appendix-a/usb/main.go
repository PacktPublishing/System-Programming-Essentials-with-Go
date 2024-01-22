package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
	"unicode"

	"github.com/godbus/dbus/v5"
)

// When the usb device is detected we are take an action
// in this example we are organizing the files in the usb
// and send a signal to be listened by the notification daemon

type Notification struct {
	AppName       string
	ReplacesID    uint32
	AppIcon       string
	Summary       string
	Body          string
	Actions       []string
	Hints         map[string]dbus.Variant
	ExpireTimeout int32
}

func NewDefaultNotification(deviceName string) Notification {
	hints := map[string]dbus.Variant{
		"urgency": dbus.MakeVariant(uint8(2)), // Use 2 for critical priority
	}
	return Notification{
		AppName:       "Organizer",
		ReplacesID:    uint32(0),
		AppIcon:       "view-refresh",
		Summary:       "Organizer is done!",
		Body:          fmt.Sprintf("The files at %s were successfully organized.", deviceName),
		Actions:       []string{},
		Hints:         hints,
		ExpireTimeout: int32(5000),
	}
}

func main() {
	conn, err := dbus.SystemBus()
	if err != nil {
		fmt.Sprintf("Failed to connect to system bus: %v\n", err)
		return
	}
	defer conn.Close()
	// Create a channel to receive signals
	ch := make(chan *dbus.Signal)
	conn.Signal(ch)

	// Add a match to listen for signals from UDisks2
	matchStr := "type='signal',sender='org.freedesktop.UDisks2',interface='org.freedesktop.DBus.ObjectManager',path='/org/freedesktop/UDisks2'"
	call := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, matchStr)
	if call.Err != nil {
		fmt.Sprintf("Failed to add match: %v\n", call.Err)
		return
	}
	notificationsCh := make(chan string)
	go func() {
		var err error
		for s := range notificationsCh {
			n := NewDefaultNotification(s)
			if err = n.send(); err != nil {
				fmt.Printf("Error sending notification: %v\n", err)
			}
		}
	}()

	for signal := range ch {
		if signal.Name == "org.freedesktop.DBus.ObjectManager.InterfacesAdded" {
			path := signal.Body[0].(dbus.ObjectPath)
			if strings.HasPrefix(string(path), "/org/freedesktop/UDisks2/block_devices/") {
				deviceObj := conn.Object("org.freedesktop.UDisks2", path)
				deviceProps := deviceObj.Call("org.freedesktop.DBus.Properties.Get", 0,
					"org.freedesktop.UDisks2.Block", "Device")

				if deviceProps.Err != nil {
					fmt.Println("Error fetching mount point:", deviceProps.Err)
					continue
				}

				mps := deviceProps.Body[0].(dbus.Variant)
				mp := removeTrailingNullChars(fmt.Sprintf("%s", mps.Value()))

				if isPartition(mp) {
					partitions, err := mountPoints([]string{mp})
					if err != nil {
						fmt.Println(err)
						continue
					}
					for _, partition := range partitions {
						partition = removeTrailingNullChars(partition)
						files, err := organizeFiles([]string{partition})
						if err != nil {
							fmt.Println(err)
							continue
						}
						for _, file := range files {
							fmt.Println(file)
						}
						notificationsCh <- partition
					}

				}
			}
		}
	}
}

func removeTrailingNullChars(s string) string {
	// Remove any trailing null characters
	for len(s) > 0 && s[len(s)-1] == '\x00' {
		s = s[:len(s)-1]
	}
	return s
}

func isPartition(s string) bool {
	if len(s) == 0 {
		return false
	}

	lastChar := s[len(s)-1]
	return unicode.IsDigit(rune(lastChar))
}

func mountPoints(deviceNames []string) ([]string, error) {
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to system bus: %v", err)
	}
	defer conn.Close()

	var mps []string
	for _, deviceName := range deviceNames {
		mps1, err := retry(5, 200*time.Millisecond, deviceName, conn, fetchMountPoints)
		if err != nil {
			return nil, err
		}
		mps = append(mps1, mps...)
	}

	return mps, nil
}

func fetchMountPoints(deviceName string, conn *dbus.Conn) ([]string, error) {
	var mountPoints []string

	partition := path.Base(deviceName)
	objPath := path.Join("/org/freedesktop/UDisks2/block_devices", partition)
	obj := conn.Object("org.freedesktop.UDisks2", dbus.ObjectPath(objPath))
	var result map[string]dbus.Variant

	err := obj.Call("org.freedesktop.DBus.Properties.GetAll", 0, "org.freedesktop.UDisks2.Filesystem").Store(&result)
	if err != nil {
		return nil, fmt.Errorf("failed to call method: %v", err)
	}

	if mountPointsVariant, exists := result["MountPoints"]; exists {
		mountPointsValue := mountPointsVariant.Value().([][]byte)
		if len(mountPointsValue) == 0 {
			return nil, fmt.Errorf("no mount points found")
		}
		for _, mp := range mountPointsValue {
			mountPoints = append(mountPoints, string(mp))
		}
	}
	return mountPoints, nil
}

// retry is a generic function that performs retries for a given function.
func retry(maxRetries int, interval time.Duration, deviceName string, conn *dbus.Conn, fn func(deviceName string, conn *dbus.Conn) ([]string, error)) ([]string, error) {
	var slice []string
	var err error
	for retries := 0; retries < maxRetries; retries++ {
		slice, err = fn(deviceName, conn)
		if err == nil {
			return slice, nil
		}
		fmt.Printf("Attempt %d failed with error: %v\n", retries+1, err)
		time.Sleep(interval)
	}
	return nil, err
}

func (n *Notification) send() error {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Errorf("failed to connect to session bus: %v", err)
	}
	defer conn.Close()

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")

	call := obj.Call("org.freedesktop.Notifications.Notify", 0,
		n.AppName, n.ReplacesID,
		n.AppIcon, n.Summary,
		n.Body, n.Actions,
		n.Hints, n.ExpireTimeout)

	if call.Err != nil {
		return fmt.Errorf("error: %v", call.Err)
	}

	return nil
}

func organizeFiles(paths []string) ([]string, error) {
	var err error
	events := make([]string, 0)
	for _, rootPath := range paths {
		err := filepath.WalkDir(rootPath, func(path string, dir os.DirEntry, err error) error {
			if err != nil {
				return err
			}
			// Check if the current dir is the rootPath or a subdirectory
			if dir.IsDir() && rootPath != path {
				return filepath.SkipDir
			}
			if !dir.IsDir() {
				ext := filepath.Ext(path)
				if len(ext) != 0 {
					destDir := filepath.Join(filepath.Dir(path), ext[1:]) // Remove the leading dot from the extension
					destPath := filepath.Join(destDir, dir.Name())

					// Create the destination directory if it doesn't exist
					if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
						return err
					}

					// Move the file to the destination
					if err := os.Rename(path, destPath); err != nil {
						return err
					}
					events = append(events, fmt.Sprintf("Moved %s to %s\n", path, destPath))
				}
			}
			return nil
		})

		if err != nil {
			fmt.Printf("Error walking the path %v: %v\n", rootPath, err)
		}
	}
	return events, err
}

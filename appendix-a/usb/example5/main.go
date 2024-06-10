package main

import (
	"fmt"
	"path"

	"github.com/godbus/dbus/v5"
)

func main() {
	points, err := mountPoints([]string{"sdc1"})
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Mount point:", points)
}

func mountPoints(deviceNames []string) ([]string, error) {
	conn, err := dbus.ConnectSystemBus()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to system bus: %v", err)
	}
	defer conn.Close()

	var mountPoints []string

	for _, deviceName := range deviceNames {
		objPath := path.Join("/org/freedesktop/UDisks2/block_devices", deviceName)
		obj := conn.Object("org.freedesktop.UDisks2", dbus.ObjectPath(objPath))
		var result map[string]dbus.Variant

		err = obj.Call("org.freedesktop.DBus.Properties.GetAll", 0, "org.freedesktop.UDisks2.Filesystem").Store(&result)
		if err != nil {
			return nil, fmt.Errorf("failed to call method: %v", err)
		}

		if mountPointsVariant, exists := result["MountPoints"]; exists {
			mountPointsValue := mountPointsVariant.Value().([][]byte)
			for _, mp := range mountPointsValue {
				mountPoints = append(mountPoints, string(mp))
			}
		}
	}

	if len(mountPoints) == 0 {
		return nil, fmt.Errorf("no mount points found")
	}

	return mountPoints, nil
}

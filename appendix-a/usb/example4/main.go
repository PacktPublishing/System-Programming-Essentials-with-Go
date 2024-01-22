package main

import (
	"fmt"
	"strings"

	"github.com/godbus/dbus/v5"
)

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

				mountPoints := deviceProps.Body[0].(dbus.Variant)
				fmt.Println(fmt.Sprintf("%s", mountPoints.Value()))
			}
		}
	}
}

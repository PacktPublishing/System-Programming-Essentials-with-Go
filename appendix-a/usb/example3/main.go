package main

import (
	"fmt"

	"github.com/godbus/dbus/v5"
)

func main() {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		fmt.Errorf("failed to connect to session bus: %v", err)
	}
	defer conn.Close()

	obj := conn.Object("org.freedesktop.Notifications", "/org/freedesktop/Notifications")
	appName := "Organizer"
	replacesID := uint32(0)
	appIcon := "view-refresh"
	summary := "Organizer is done!"
	body := fmt.Sprintf("The files at %s were successfully organized.", "/dev/sdc")
	actions := []string{}
	hints := map[string]dbus.Variant{}
	expireTimeout := int32(5000)
	call := obj.Call("org.freedesktop.Notifications.Notify", 0, appName, replacesID,
		appIcon, summary, body, actions,
		hints, expireTimeout)
	if call.Err != nil {
		fmt.Sprintf("Error: %v", call.Err)
		return
	}
}

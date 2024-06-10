package main

import (
	"fmt"
	"os/exec"

	"github.com/muka/go-bluetooth/api"
)

func main() {
	adapter, err := api.GetDefaultAdapter()
	if err != nil {
		fmt.Printf("Failed to find default adapter: %s\n", err)
	}

	err = adapter.StartDiscovery()
	if err != nil {
		fmt.Printf("Failed to start discovery: %s\n", err)
	}

	devices, err := adapter.GetDevices()
	if err != nil {
		fmt.Printf("Failed to get devices: %s\n", err)
	}

	for _, device := range devices {
		info, err := device.GetProperties()
		if err != nil {
			fmt.Printf("Failed to get properties: %s\n", err)
			continue
		}

		if info.RSSI < -70 && info.Name == "Galaxy Watch Active2(207D) LE" {
			err := lockScreen()
			if err != nil {
				fmt.Printf("Failed to lock screen: %s\n", err)
				continue
			}
		}
	}
}

func lockScreen() error {
	_, err := exec.Command("xdg-screensaver", "lock").Output()
	if err != nil {
		return err
	}
	return nil
}

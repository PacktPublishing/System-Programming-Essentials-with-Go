package main

import (
	"fmt"

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
		fmt.Println(info.Name, info.Address, info.RSSI)
	}
}

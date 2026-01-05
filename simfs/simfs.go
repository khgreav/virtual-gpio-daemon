package simfs

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
)

const (
	BASE_PATH     = "/sys/kernel/config"
	GPIO_SIM_PATH = BASE_PATH + "/gpio-sim"
)

func CheckInit() error {
	_, err := os.Stat(BASE_PATH)
	if err != nil {
		return fmt.Errorf("Configfs is not mounted at %s: %w", BASE_PATH, err)
	}

	_, err = os.Stat(GPIO_SIM_PATH)
	if err != nil {
		return fmt.Errorf("gpio-sim kernel module not loaded: %w", err)
	}

	err = unix.Access(GPIO_SIM_PATH, unix.W_OK)
	if err != nil {
		return fmt.Errorf("Insufficient permissions for %s: %w", GPIO_SIM_PATH, err)
	}

	return nil
}

func Cleanup(devices []Device) error {
	for _, device := range devices {
		err := device.Delete()
		if err != nil {
			return fmt.Errorf("Failed to clean up gpio-sim device %s: %s", device.DevName, err.Error())
		}
	}
	return nil
}

func Initialize(devices []Device) error {
	for _, device := range devices {
		err := device.Create()
		if err != nil {
			return fmt.Errorf("Failed to create device %s: %s", device.DevName, err.Error())
		}
	}
	return nil
}

package simfs

import (
	"fmt"
	"os"

	"golang.org/x/sys/unix"
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

func Cleanup() error {
	err := os.RemoveAll(GPIO_SIM_PATH)
	if err != nil {
		return fmt.Errorf("Failed to clean up gpio-sim configfs entries: %w", err)
	}
	return nil
}

func Initialize(devices []Device) error {
	for idx, device := range devices {
		err := device.Create()
		if err != nil {
			return fmt.Errorf("Failed to create %dth device: %s", idx, err.Error())
		}
	}
	return nil
}

func (device Device) Create() error {
	return nil
}

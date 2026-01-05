package simfs

import (
	"fmt"
	"os"
)

const (
	DEVICE_NAME = "/dev_name"
	DEVICE_LIVE = "/live"
)

type Device struct {
	DevName string `yaml:"dev_name"`
	Live    bool   `yaml:"live"`
	Banks   []Bank `yaml:"banks"`
}

func (device Device) GetLiveValue() string {
	if device.Live {
		return "1"
	}
	return "0"
}

func (device Device) Create() error {
	err := os.MkdirAll(GPIO_SIM_PATH+"/"+device.DevName, 0755)
	if err != nil {
		return fmt.Errorf("Failed to create device directory %s: %s", device.DevName, err.Error())
	}
	for _, bank := range device.Banks {
		err = bank.Create(device.DevName)
		if err != nil {
			return fmt.Errorf("Failed to create bank %s: %s", bank.ChipName, err.Error())
		}
	}
	err = os.WriteFile(
		GPIO_SIM_PATH+"/"+device.DevName+DEVICE_LIVE,
		[]byte(device.GetLiveValue()),
		0644,
	)
	if err != nil {
		return fmt.Errorf("Failed to set device live: %s", err.Error())
	}
	return nil
}

func (device Device) Delete() error {
	os.WriteFile(
		GPIO_SIM_PATH+"/"+device.DevName+DEVICE_LIVE,
		[]byte("0"),
		0644,
	)
	for _, bank := range device.Banks {
		bank.Delete(device.DevName)
	}
	os.Remove(GPIO_SIM_PATH + "/" + device.DevName)
	return nil
}

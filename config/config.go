package config

import (
	"fmt"
	"os"

	"github.com/khgreav/virtual-gpio-daemon/simfs"
	"gopkg.in/yaml.v3"
)

func LoadConfig(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("Failed to read config file %s: %w", path, err)
	}
	return data, nil
}

func Parse(data []byte) ([]simfs.Device, error) {
	var devices []simfs.Device

	err := yaml.Unmarshal(data, &devices)
	if err != nil {
		return nil, fmt.Errorf("Failed to parse config data: %w", err)
	}

	for idx, device := range devices {
		err := device.Validate()
		if err != nil {
			return nil, fmt.Errorf("Failed to validate %dth device: %s", idx, err.Error())
		}
	}

	return devices, nil
}

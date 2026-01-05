package simfs

import (
	"fmt"
	"os"
	"strconv"
)

const (
	BANK_CHIP     = "/chip_name"
	BANK_NUMLINES = "/num_lines"
)

type Bank struct {
	ChipName string `yaml:"chip_name"`
	NumLines int    `yaml:"num_lines"`
	Lines    []Line `yaml:"lines"`
}

func (bank Bank) Create(device string) error {
	err := os.MkdirAll(GPIO_SIM_PATH+"/"+device+"/"+bank.ChipName, 0755)
	if err != nil {
		return fmt.Errorf("Failed to create bank directory %s: %s", bank.ChipName, err.Error())
	}
	err = os.WriteFile(
		GPIO_SIM_PATH+"/"+device+"/"+bank.ChipName+BANK_NUMLINES,
		[]byte(strconv.Itoa(bank.NumLines)),
		0644,
	)
	if err != nil {
		return fmt.Errorf("Failed to bank num lines: %s", err.Error())
	}
	for _, line := range bank.Lines {
		err = line.Create(device, bank.ChipName)
		if err != nil {
			return fmt.Errorf("Failed to create line %s: %s", line.Name, err.Error())
		}
	}
	return nil
}

func (bank Bank) Delete(device string) error {
	for _, line := range bank.Lines {
		line.Delete(device, bank.ChipName)
	}
	os.Remove(GPIO_SIM_PATH + "/" + device + "/" + bank.ChipName)
	return nil
}

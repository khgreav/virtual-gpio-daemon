package simfs

import "fmt"

func (device Device) Validate() error {
	if device.DevName == "" {
		return fmt.Errorf("Device name cannot be empty.")
	}
	for idx, bank := range device.Banks {
		err := bank.Validate()
		if err != nil {
			return fmt.Errorf("Failed to validate %dth bank: %s", idx, err.Error())
		}
	}
	return nil
}

func (bank Bank) Validate() error {
	if bank.ChipName == "" {
		return fmt.Errorf("Bank chip name cannot be empty.")
	}
	if bank.NumLines <= 0 {
		return fmt.Errorf("Bank must have at least one line.")
	}
	if len(bank.Lines) != bank.NumLines {
		return fmt.Errorf("Bank NumLines (%d) does not match number of defined lines (%d).", bank.NumLines, len(bank.Lines))
	}
	for idx, line := range bank.Lines {
		err := line.Validate()
		if err != nil {
			return fmt.Errorf("Failed to validate %dth line: %s", idx, err.Error())
		}
	}
	return nil
}

func (line Line) Validate() error {
	if line.Name == "" {
		return fmt.Errorf("Line name cannot be empty.")
	}
	return nil
}

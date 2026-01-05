package simfs

import (
	"fmt"
	"os"
)

const (
	LINE_NAME  = "/name"
	LINE_VALID = "/valid"
)

type Line struct {
	Name  string `yaml:"name"`
	Valid bool   `yaml:"valid"`
}

func (line Line) GetValidValue() string {
	if line.Valid {
		return "1"
	}
	return "0"
}

func (line Line) Create(device string, bank string) error {
	err := os.MkdirAll(GPIO_SIM_PATH+"/"+device+"/"+bank+"/"+line.Name, 0755)
	if err != nil {
		return fmt.Errorf("Failed to create line directory %s: %s", line.Name, err.Error())
	}
	_, err = os.Stat(GPIO_SIM_PATH + "/" + device + "/" + bank + "/" + line.Name + LINE_VALID)
	if err != nil {
		return nil
	}
	err = os.WriteFile(
		GPIO_SIM_PATH+"/"+device+"/"+bank+"/"+line.Name+LINE_VALID,
		[]byte(line.GetValidValue()),
		0644,
	)
	if err != nil {
		return fmt.Errorf("Failed to line valid: %s", err.Error())
	}
	return nil
}

func (line Line) Delete(device string, bank string) error {
	os.Remove(GPIO_SIM_PATH + "/" + device + "/" + bank + "/" + line.Name)
	return nil
}

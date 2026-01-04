package simfs

type Device struct {
	DevName string `yaml:"dev_name"`
	Live    bool   `yaml:"live"`
	Banks   []Bank `yaml:"banks"`
}

type Bank struct {
	ChipName string `yaml:"chip_name"`
	NumLines int    `yaml:"num_lines"`
	Lines    []Line `yaml:"lines"`
}

type Line struct {
	Name  string `yaml:"name"`
	Valid bool   `yaml:"valid"`
}

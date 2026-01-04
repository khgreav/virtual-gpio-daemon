package simfs

const (
	BASE_PATH     = "/sys/kernel/config"
	GPIO_SIM_PATH = BASE_PATH + "/gpio-sim"

	DEVICE_NAME = "/dev_name"
	DEVICE_LIVE = "/live"

	BANK_PATH     = "/bank%d"
	BANK_CHIP     = "/chip_name"
	BANK_NUMLINES = "/num_lines"

	LINE_PATH  = "/line%d"
	LINE_NAME  = "/name"
	LINE_VALID = "/valid"
)

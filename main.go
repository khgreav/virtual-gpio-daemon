package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/khgreav/virtual-gpio-daemon/config"
	"github.com/khgreav/virtual-gpio-daemon/simfs"
)

type Args struct {
	ConfigFile string
}

func parseArgs() Args {
	var args Args

	flag.StringVar(&args.ConfigFile, "config", "", "Path to yaml configuration file")
	flag.Parse()

	if args.ConfigFile == "" {
		fmt.Fprintf(os.Stderr, "--config option is required.\n")
		os.Exit(1)
	}

	if !strings.HasSuffix(args.ConfigFile, ".yaml") && !strings.HasSuffix(args.ConfigFile, ".yml") {
		fmt.Fprintf(os.Stderr, "--config file must point to a .yaml or .yml file.\n")
		os.Exit(1)
	}

	return args
}

func main() {
	args := parseArgs()
	data, err := config.LoadConfig(args.ConfigFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}
	devices, err := config.Parse(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}
	err = simfs.CheckInit()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}
	err = simfs.Initialize(devices)
	if err != nil {
		simfs.Cleanup()
		fmt.Fprintf(os.Stderr, "%s", err.Error())
		os.Exit(1)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals,
		syscall.SIGINT,  // Interrupt
		syscall.SIGTERM, // Termination
	)

	for {
		sig := <-signals
		switch sig {
		case syscall.SIGINT, syscall.SIGTERM:
			err := simfs.Cleanup()
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s", err.Error())
				os.Exit(1)
			}
			os.Exit(0)
		}
	}
}

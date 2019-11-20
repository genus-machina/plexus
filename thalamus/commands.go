package main

import (
	"errors"
	"log"
	"os"

	"github.com/genus-machina/plexus/zygote"
)

func activate(logger *log.Logger, state *store, args []string) error {
	if l := len(args); l < 1 {
		return errors.New("activate requires a device name.")
	}

	if state.Devices == nil {
		return errors.New("No devices have been loaded.")
	}

	return state.Devices.DeviceBus.Activate(args[0])
}

func deactivate(logger *log.Logger, state *store, args []string) error {
	if l := len(args); l < 1 {
		return errors.New("deactivate requires a device name.")
	}

	if state.Devices == nil {
		return errors.New("No devices have been loaded.")
	}

	return state.Devices.DeviceBus.Deactivate(args[0])
}

func load(logger *log.Logger, state *store, args []string) error {
	if l := len(args); l < 1 {
		return errors.New("load requires a path.")
	}

	devices, err := zygote.LoadJSON(args[0], logger)
	if err == nil {
		if state.Devices != nil {
			state.Devices.DeviceBus.Halt()
		}
		state.Devices = devices
	}
	return err
}

func quit(logger *log.Logger, state *store, args []string) error {
	if state.Devices != nil {
		state.Devices.DeviceBus.Halt()
	}

	logger.Println("Goodbye!")
	os.Exit(0)
	return nil
}

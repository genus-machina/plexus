package main

import (
	"errors"
	"log"
	"os"

	"github.com/genus-machina/plexus/amygdala"
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

func clear(logger *log.Logger, state *store, args []string) error {
	if state.Devices == nil {
		return errors.New("No devices have been loaded.")
	}

	if state.Devices.Screen == nil {
		return errors.New("No screen has been configured.")
	}

	return state.Devices.Screen.Clear()
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

func display(logger *log.Logger, state *store, args []string) error {
	if l := len(args); l < 1 {
		return errors.New("display requires a file path.")
	}

	if state.Devices == nil {
		return errors.New("No devices have been loaded.")
	}

	if state.Devices.Screen == nil {
		return errors.New("No screen has been configured.")
	}

	var err error
	var png amygdala.Widget
	if png, err = amygdala.NewPNG(args[0]); err == nil {
		err = state.Devices.Screen.Render(png)
	}
	return err
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

func publish(logger *log.Logger, state *store, args []string) error {
	if l := len(args); l < 1 {
		return errors.New("publish requires a topic.")
	}

	if l := len(args); l < 2 {
		return errors.New("publish requires a message.")
	}

	if state.Devices == nil {
		return errors.New("No devices have been loaded.")
	}

	if state.Devices.Synapse == nil {
		return errors.New("A synapse has not been configured.")
	}

	topic := args[0]
	message := []byte(args[1])
	return state.Devices.Synapse.Publish(message, topic)
}

func quit(logger *log.Logger, state *store, args []string) error {
	state.Devices.Halt()
	logger.Println("Goodbye!")
	os.Exit(0)
	return nil
}

func rotate(logger *log.Logger, state *store, args []string) error {
	if l := len(args); l < 1 {
		return errors.New("display requires a rotation angle.")
	}

	if state.Devices == nil {
		return errors.New("No devices have been loaded.")
	}

	if state.Devices.Screen == nil {
		return errors.New("No screen has been configured.")
	}

	switch args[0] {
	default:
	case "0":
		state.Devices.Screen.Rotate(amygdala.IMAGE_ROTATE_0)
	case "90":
		state.Devices.Screen.Rotate(amygdala.IMAGE_ROTATE_90)
	case "180":
		state.Devices.Screen.Rotate(amygdala.IMAGE_ROTATE_180)
	case "270":
		state.Devices.Screen.Rotate(amygdala.IMAGE_ROTATE_270)
	}

	return nil
}

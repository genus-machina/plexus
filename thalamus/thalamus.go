package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/genus-machina/plexus/hippocampus"
	"github.com/genus-machina/plexus/zygote"
)

type store struct {
	Devices *zygote.System
}

func execute(logger *log.Logger, command []string, state *store) error {
	if len(command) == 0 {
		return nil
	}

	var err error
	name := command[0]
	args := command[1:]

	switch name {
	case "activate":
		err = activate(logger, state, args)
	case "deactivate":
		err = deactivate(logger, state, args)
	case "publish":
		err = publish(logger, state, args)
	case "load":
		err = load(logger, state, args)
	case "quit":
		quit(logger, state, args)
	default:
		return fmt.Errorf("Unrecognized command '%s'.", name)
	}

	return err
}

func parse(input string) ([]string, error) {
	parser := bufio.NewScanner(strings.NewReader(input))
	parser.Split(splitToken)
	tokens := make([]string, 0)

	for parser.Scan() {
		tokens = append(tokens, parser.Text())
	}

	return tokens, parser.Err()
}

func main() {
	logger := hippocampus.NewLogger("thalamus")
	reader := bufio.NewReader(os.Stdin)
	state := new(store)

	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')

		if err == io.EOF {
			quit(logger, state, nil)
		} else if err != nil {
			logger.Fatalf("error: %s", err.Error())
		}

		command, err := parse(input)
		if err != nil {
			logger.Printf("error: %s", err.Error())
		}

		if err := execute(logger, command, state); err != nil {
			logger.Printf("error: %s", err.Error())
		}
	}

	quit(logger, state, nil)
}

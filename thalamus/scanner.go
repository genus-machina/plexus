package main

import (
	"bytes"
	"errors"
	"io"
	"unicode"
)

const (
	pending = iota
	simple
	compound
	complete
)

func processRune(state int, current rune, output *bytes.Buffer) int {
	switch state {
	case pending:
		if unicode.IsSpace(current) {
			return pending
		} else if current == '"' {
			return compound
		} else {
			output.WriteRune(current)
			return simple
		}
	case simple:
		if unicode.IsSpace(current) {
			return complete
		} else {
			output.WriteRune(current)
			return simple
		}
	case compound:
		if current == '"' {
			return complete
		} else {
			output.WriteRune(current)
			return compound
		}
	default:
		return complete
	}
}

func splitToken(data []byte, atEOF bool) (int, []byte, error) {
	advance := 0
	input := bytes.NewBuffer(data)
	output := new(bytes.Buffer)
	state := pending

	for r, size, err := input.ReadRune(); state != complete; r, size, err = input.ReadRune() {
		if err == io.EOF {
			break
		} else if err != nil {
			return 0, nil, err
		}

		state = processRune(state, r, output)
		advance += size
	}

	if atEOF {
		if state == pending {
			return 0, nil, nil
		}
		if state == compound {
			return 0, nil, errors.New("Unterminated \".")
		}
	} else {
		return 0, nil, nil
	}

	return advance, output.Bytes(), nil
}

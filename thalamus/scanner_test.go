package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestSplitTokenEmpty(t *testing.T) {
	data := []byte("")
	_, token, err := splitToken(data, true)

	if err != nil {
		t.Errorf("split error: %s", err.Error())
	}

	if token != nil {
		t.Errorf("expected nil token bus was %s", token)
	}
}

func TestSplitTokenWhiteSpace(t *testing.T) {
	data := []byte(" ")
	_, token, err := splitToken(data, true)

	if err != nil {
		t.Errorf("split error: %s", err.Error())
	}

	if token != nil {
		t.Errorf("expect nil token but was %s", token)
	}
}

func TestSplitTokenSimpleToken(t *testing.T) {
	data := []byte("token")
	advance, token, err := splitToken(data, true)

	if err != nil {
		t.Errorf("split error: %s", err.Error())
	}

	if l := len(data); advance != l {
		t.Errorf("expected advance of %d but was %d", l, advance)
	}

	if !bytes.Equal(token, data) {
		t.Errorf("expect token '%s' but was '%s'", data, token)
	}
}

func TestSplitTokenPaddedSimpleToken(t *testing.T) {
	data := []byte("  token  ")
	advance, token, err := splitToken(data, true)

	if err != nil {
		t.Errorf("split error: %s", err.Error())
	}

	if expected := len(data) - 1; advance != expected {
		t.Errorf("expected advance of %d but was %d", expected, advance)
	}

	if expected := []byte("token"); !bytes.Equal(token, expected) {
		t.Errorf("expected token '%s' but was '%s'", expected, token)
	}
}

func TestSplitTokenIncompleteSimpleToken(t *testing.T) {
	data := []byte("token")
	advance, token, err := splitToken(data, false)

	if err != nil {
		t.Errorf("split error: %s", err.Error())
	}

	if advance != 0 {
		t.Errorf("expected advance of 0 but was %d", advance)
	}

	if token != nil {
		t.Errorf("expected nil token but was %s", token)
	}
}

func TestSplitTokenCompoundToken(t *testing.T) {
	data := []byte("\"hello world\"")
	advance, token, err := splitToken(data, true)

	if err != nil {
		t.Errorf("split error: %s", err.Error())
	}

	if l := len(data); advance != l {
		t.Errorf("expected advance of %d but was %d", l, advance)
	}

	if expected := []byte("hello world"); !bytes.Equal(token, expected) {
		t.Errorf("expected token '%s' but was '%s'", expected, token)
	}
}

func TestSplitTokenIncompleteCompoundToken(t *testing.T) {
	data := []byte("\"hello world")
	advance, token, err := splitToken(data, false)

	if err != nil {
		t.Errorf("split error: %s", err.Error())
	}

	if advance != 0 {
		t.Errorf("expected advance of 0 but was %d", advance)
	}

	if token != nil {
		t.Errorf("expected nil token but was %s", token)
	}
}

func TestSplitTokenUnterminatedCompoundToken(t *testing.T) {
	data := []byte("\"hello world")
	_, _, err := splitToken(data, true)

	if message := err.Error(); !strings.Contains(message, "Unterminated") {
		t.Errorf("unexpected error: %s", message)
	}
}

func TestSplitTokenPaddedCompoundToken(t *testing.T) {
	data := []byte("  \"compound token\"  ")
	advance, token, err := splitToken(data, true)

	if err != nil {
		t.Errorf("split error: %s", err.Error())
	}

	if expected := len(data) - 2; advance != expected {
		t.Errorf("expected advance of %d but was %d", expected, advance)
	}

	if expected := []byte("compound token"); !bytes.Equal(token, expected) {
		t.Errorf("expected token '%s' but was '%s'", expected, token)
	}
}

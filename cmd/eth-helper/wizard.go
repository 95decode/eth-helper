package main

import (
	"fmt"
	"strings"

	"github.com/95decode/eth-helper/console/prompt"
	"github.com/peterh/liner"
)

type wizard struct {
	flag string // additional config
	// lock sync.Mutex // Lock to protect configs during concurrent service discovery
}

func (w *wizard) read() string {
	text := promptInput("> ")
	return strings.TrimSpace(text)
}

func (w *wizard) readPassword() string {
	text := promptPassword("> ")
	return strings.TrimSpace(text)
}

func (w *wizard) confirm() bool {
	input := promptInput("> [y/n] > ")
	if len(input) > 0 && strings.EqualFold(input[:1], "y") {
		return true
	}
	return false
}

func promptInput(p string) string {
	for {
		text, err := prompt.Stdin.PromptInput(p)
		if err != nil {
			if err != liner.ErrPromptAborted {
				fmt.Println("Failed to read user input", "err", err)
			}
		} else {
			return text
		}
	}
}

func promptPassword(p string) string {
	for {
		text, err := prompt.Stdin.PromptPassword(p)
		if err != nil {
			if err != liner.ErrPromptAborted {
				fmt.Println("Failed to read user input", "err", err)
			}
		} else {
			return text
		}
	}
}

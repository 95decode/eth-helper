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

func (w *wizard) confirm() bool {
	text := promptInput("> ")
	switch text {
	case "y":
		return true
	case "n":
		return false
	default:
		return false
	}
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

package main

import (
	"fmt"
)

// makeWizard creates and returns a new puppeth wizard.
func makeWizard(flag string) *wizard {
	return &wizard{
		flag: flag,
	}
}

func (w *wizard) run() {
	fmt.Println("+-----------------------------------------------------------+")
	fmt.Println("| test 													 |")
	fmt.Println("+-----------------------------------------------------------+")
	fmt.Println()

	for {
		fmt.Println()
		fmt.Println("What would you like to do? (default = 1)")
		fmt.Println(" 1. foo")
		fmt.Println(" 2. bar")
		fmt.Println(" 3. flag test")

		choice := w.read()
		switch {
		case choice == "" || choice == "1":
			fmt.Println("fooooooo")
		case choice == "2":
			fmt.Println("baaaaaar")
		case choice == "3":
			fmt.Println(w.flag)
		default:
			fmt.Println("That's not something I can do")
		}
	}
}

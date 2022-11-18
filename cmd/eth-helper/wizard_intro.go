package main

import (
	"fmt"

	"github.com/95decode/eth-helper/utils/key"
)

func makeWizard(flag string) *wizard {
	return &wizard{
		flag: flag,
	}
}

func (w *wizard) run() {
	fmt.Println("+-----------------------------------------------------------+")
	fmt.Println("|------------------------ETH  HELPER------------------------|")
	fmt.Println("+-----------------------------------------------------------+")
	fmt.Println()

	for {
		fmt.Println()
		fmt.Println("What would you like to do? (default = 1)")
		fmt.Println(" 1. Create account")
		fmt.Println(" 2. bar")
		fmt.Println(" 3. flag test")

		choice := w.read()
		switch {
		case choice == "" || choice == "1":
			fmt.Println()
			fmt.Println("Encrypt")

			encrypted := w.confirm()
			passphrase := ""
			_, ks := key.GenerateKey(encrypted)

			if encrypted {
				fmt.Println()
				fmt.Println("passphrase")
				passphrase = w.readPassword()
			}

			a, err := ks.NewAccount(passphrase)
			if err != nil {
				fmt.Println("error")
			}

			fmt.Println()
			fmt.Println("Address : ", a.Address)
		case choice == "2":
			fmt.Println("222")
		case choice == "3":
			fmt.Println(w.flag)
		default:
			fmt.Println("That's not something I can do")
		}
	}
}

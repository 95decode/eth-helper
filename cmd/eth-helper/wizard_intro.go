package main

import (
	"fmt"

	keystore "github.com/ethereum/go-ethereum/accounts/keystore"
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
		fmt.Println(" 1. Create account")
		fmt.Println(" 2. bar")
		fmt.Println(" 3. flag test")

		choice := w.read()
		switch {
		case choice == "" || choice == "1":
			fmt.Println("fooooooo")

			_, ks := tmp(true)

			a, err := ks.NewAccount("foo")
			if err != nil {
				fmt.Println("error 1")
			}

			fmt.Println(a.Address)
		case choice == "2":
			fmt.Println("baaaaaar")
		case choice == "3":
			fmt.Println(w.flag)
		default:
			fmt.Println("That's not something I can do")
		}
	}
}

func tmp(encrypted bool) (string, *keystore.KeyStore) {
	d := "./"
	newKs := keystore.NewPlaintextKeyStore
	if encrypted {
		newKs = func(kd string) *keystore.KeyStore { return keystore.NewKeyStore(kd, 2, 1) }
	}
	return d, newKs(d)
}

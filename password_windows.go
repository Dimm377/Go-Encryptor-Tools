//go:build windows
// +build windows

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func getPassword() []byte {
	fmt.Print("Enter password: ")

	//(hiding characters)
	password, err := term.ReadPassword(int(os.Stdin.Fd()))

	if err != nil {
		fmt.Println() 
		fmt.Print("Warning: Could not hide password input. Enter password: ")
		reader := bufio.NewReader(os.Stdin)
		passwordStr, _ := reader.ReadString('\n')
		passwordStr = strings.TrimSpace(passwordStr) 
		fmt.Print("Confirm password: ")
		confirmStr, _ := reader.ReadString('\n')
		confirmStr = strings.TrimSpace(confirmStr) 

		if passwordStr != confirmStr {
			fmt.Println("\nPasswords do not match, please try again")
			return getPassword()
		}

		return []byte(passwordStr)
	} else {
		fmt.Println() // Add newline after password input
		fmt.Print("Confirm password: ")
		confirm, err := term.ReadPassword(int(os.Stdin.Fd()))

		if err != nil {
			// If confirmation fails, fall back to visible input
			fmt.Println()
			fmt.Print("Warning: Could not hide confirmation input. Enter password: ")
			reader := bufio.NewReader(os.Stdin)
			passwordStr, _ := reader.ReadString('\n')
			passwordStr = strings.TrimSpace(passwordStr)

			fmt.Print("Confirm password: ")
			confirmStr, _ := reader.ReadString('\n')
			confirmStr = strings.TrimSpace(confirmStr)

			if passwordStr != confirmStr {
				fmt.Println("\nPasswords do not match, please try again")
				return getPassword()
			}

			return []byte(passwordStr)
		} else {
			fmt.Println() // Add newline after confirmation
			if !validatePassword(password, confirm) {
				fmt.Println("\nPasswords do not match, please try again")
				return getPassword()
			}
			return password
		}
	}
}

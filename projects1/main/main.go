package main

import (
	"bufio"
	"fmt"
	"os"
	"projects1/logIn"
	"projects1/userEntryOptions"
	"strings"
)

func main() {
	for {
		fmt.Println("Select an option:")
		fmt.Println("1. Admin")
		fmt.Println("2. User")
		fmt.Println("3. Exit")

		// Read user input
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		// Trim any trailing newline characters
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			logIn.LogIn("admin")
		case "2":
			fmt.Println("You selected User.")
			userEntryOptions.UserEntryOptions()
		case "3":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}

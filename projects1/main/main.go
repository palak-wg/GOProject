package main

import (
	"fmt"
	"projects1/logIn"
	"projects1/signUp"
)

func main() {
	// Use a for loop to display options and ask for user choice
	for chooseOption := true; chooseOption; {
		// choiceLogSign ask choice of user for sign-up or login
		var choiceLogSign string
		fmt.Println("----------O P T I O N S:--------- \n1. SIGNUP\n2. LOGIN\n3. EXIT\nEnter your choice:")
		_, errScan := fmt.Scan(&choiceLogSign)
		if errScan != nil {
			fmt.Println("Error:", errScan)
			return
		}

		// Use a switch statement to handle different cases
		switch choiceLogSign {
		case "1", "SignUp", "signup", "SIGNUP", "1.":
			fmt.Println("--------SIGNUP-------")
			chooseOption = !signUp.SignUp()
		case "2", "login", "Login", "LOGIN", "2.":
			fmt.Println("--------LOGIN--------")
			chooseOption = !logIn.LogIn()
		case "exit", "Exit", "3", "3.", "EXIT":
			fmt.Println("--------EXIT--------")
			break
		default:
			fmt.Println("Enter valid input")
		}
	}
}

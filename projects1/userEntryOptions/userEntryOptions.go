package userEntryOptions

import (
	"fmt"
	"projects1/logIn"
	"projects1/signUp"
)

// UserEntryOptions provide user entry options like login, signup and back

func UserEntryOptions() {
	for {
		var choiceLogSign string
		fmt.Println("----------O P T I O N S:--------- \n1. SIGNUP\n2. LOGIN\n3. BACK\nEnter your choice:")
		_, errScan := fmt.Scan(&choiceLogSign)
		if errScan != nil {
			fmt.Println("Error:", errScan)
			return
		}

		// Use a switch statement to handle different cases
		switch choiceLogSign {
		case "1":
			fmt.Println("--------SIGNUP-------")
			_ = !signUp.SignUp()
		case "2":
			fmt.Println("--------LOGIN--------")
			logIn.LogIn("user")
		case "3":
			fmt.Println("--------BACK--------")
			return
		default:
			fmt.Println("Enter valid input")
		}
	}
}

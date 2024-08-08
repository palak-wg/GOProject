package user

//
//import (
//	"fmt"
//	"projects1/logIn"
//	"projects1/signUp"
//)
//
//// User provide user entry options like login, signup and back
//func Userr() {
//	// Use a for loop to display options and ask for user choice
//	for {
//		var choiceLogSign string
//		fmt.Println("----------O P T I O N S:--------- \n1. SIGNUP\n2. LOGIN\n3. BACK\nEnter your choice:")
//		_, errScan := fmt.Scan(&choiceLogSign)
//		if errScan != nil {
//			fmt.Println("Error:", errScan)
//			return
//		}
//
//		// Use a switch statement to handle different cases
//		switch choiceLogSign {
//		case "1", "SignUp", "signup", "SIGNUP", "1.":
//			fmt.Println("--------SIGNUP-------")
//			_ = !signUp.SignUp()
//		case "2", "login", "Login", "LOGIN", "2.":
//			fmt.Println("--------LOGIN--------")
//			_ = !logIn.LogIn("user")
//		case "exit", "Exit", "3", "3.", "EXIT":
//			fmt.Println("--------BACK--------")
//			break
//		default:
//			fmt.Println("Enter valid input")
//		}
//	}
//}

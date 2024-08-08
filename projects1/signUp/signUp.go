package signUp

import (
	"bufio"
	"fmt"
	"os"
	"projects1/util"
	"regexp"
	"strconv"
	"strings"
)

// SignUp asks for credentials like userId, password, age, phone number, address, and designation
func SignUp() bool {
	rd := bufio.NewReader(os.Stdin)
	var newID, newPassword, newName, phoneNumber, address, designation string

	// User ID input
	for exists := true; exists; {
		fmt.Print("Enter User ID: ")
		newID, _ = rd.ReadString('\n')
		newID = strings.TrimSpace(newID)
		if util.TakenUserId(newID) {
			fmt.Println("UserID unavailable")
			continue
		}
		exists = false
	}

	// Password input
	for strongPass := true; strongPass; {
		//fmt.Println("Password must contain:\n8 characters\nAt least 1 uppercase char\nAt least 1 lowercase char\nAt least 1 numeric char\nAt least 1 special char")
		fmt.Println("Password must contain:\n(8 characters->At least 1 uppercase char, lowercase char, 1 numeric char, 1 special char")
		fmt.Print("Enter Password: ")
		newPassword, _ = rd.ReadString('\n')
		newPassword = strings.TrimSpace(newPassword)
		if !StrongPassword(newPassword) {
			fmt.Println("Password criteria don't match")
			continue
		}
		strongPass = false
	}

	// Age input
	fmt.Print("Enter Age: ")
	ageString, _ := rd.ReadString('\n')
	ageString = strings.TrimSpace(ageString)
	age, errAge := strconv.Atoi(ageString)

	if errAge != nil {
		fmt.Println("Invalid age input:", errAge)
		return false
	}
	if age < 16 {
		fmt.Println("Ineligible User")
		return false
	}

	for {
		fmt.Print("Enter Name: ")
		newName, _ = rd.ReadString('\n')
		newName = strings.TrimSpace(newName)
		if !util.IsValidName(newName) {
			fmt.Println("Invalid name.")
			continue
		}
		break
	}

	// Phone Number input
	for {
		fmt.Print("Enter Phone Number (10 digits): ")
		phoneNumber, _ = rd.ReadString('\n')
		phoneNumber = strings.TrimSpace(phoneNumber)
		if !(util.IsValidPhoneNumber(phoneNumber)) {
			fmt.Println("Invalid phone number. It must be 10 digits.")
			continue
		}
		break
	}

	// Address input
	fmt.Print("Enter Address: ")
	address, _ = rd.ReadString('\n')
	address = strings.TrimSpace(address)

	// Designation input
	fmt.Print("Enter Designation: ")
	designation, _ = rd.ReadString('\n')
	designation = strings.TrimSpace(designation)

	// Add user to the system
	errAddUser := util.AddUser(newID, util.HashSHA256(newPassword), newName, phoneNumber, address, designation)
	if errAddUser != nil {
		fmt.Println("Error adding user:", errAddUser)
		return false
	}
	fmt.Println("Success")
	return true
}

// StrongPassword checks if the password meets the criteria
func StrongPassword(password string) bool {
	if len(password) < 8 {
		return false
	}

	hasUpper := regexp.MustCompile(`[A-Z]`).MatchString(password)
	hasLower := regexp.MustCompile(`[a-z]`).MatchString(password)
	hasDigit := regexp.MustCompile(`[0-9]`).MatchString(password)
	hasSpecial := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)

	return hasUpper && hasLower && hasDigit && hasSpecial
}

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

const FileName = "credentials.txt"

// SignUp asks for credentials like userId, password & age
func SignUp() bool {
	// Enter credentials i.e. UserName, Password & age
	rd := bufio.NewReader(os.Stdin)
	var newID, newPassword string
	for exists := true; exists; {

		fmt.Print("Enter User ID: ")
		newID, _ = rd.ReadString('\n')
		newID = strings.TrimSpace(newID)
		if TakenUserId(newID) {
			fmt.Println("UserID unavailable")
			continue
		}
		exists = false
	}
	for strongPass := true; strongPass; {
		fmt.Println("Password must contain:-\n8 characters\nAtleast 1 uppercase char\nAtleast 1 lowercase char\nAt least 1 numeric char\nAt least 1 special char")
		fmt.Print("Enter Password: ")
		newPassword, _ = rd.ReadString('\n')
		if !StrongPassword(newPassword) {
			fmt.Println("Password criteria don't match")
			continue
		}
		strongPass = false
	}
	fmt.Print("Enter Age: ")
	ageString, _ := rd.ReadString('\n')
	ageString = strings.TrimSpace(ageString)
	age, errAge := strconv.Atoi(ageString)

	// Constraints applied
	if errAge != nil {
		panic(errAge)
	}
	if age >= 16 {
		errAddUser := util.AddUser(newID, newPassword)
		if errAddUser != nil {
			return false
		}
		fmt.Println("Success")
	} else {
		fmt.Println("Ineligible User")
		return false
	}
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

func TakenUserId(uID string) bool {
	existCred := util.ReadCredFile()
	return existCred["UserID"] == uID
}

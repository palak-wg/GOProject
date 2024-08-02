package InitialLog

import (
	"bufio"
	"fmt"
	"os"
	"projects1/CheckCred"
	"strconv"
	"strings"
)

const FileName = "empty.txt"

func Takes(i int) {
	// Enter new credentials
	fmt.Print("Enter User ID: ")
	rd := bufio.NewReader(os.Stdin)
	newID, _ := rd.ReadString('\n')
	newID = strings.TrimSpace(newID)

	fmt.Print("Enter Password: ")
	newPassword, _ := rd.ReadString('\n')
	newPassword = strings.TrimSpace(newPassword)
	if i == 1 {
		TakesSignUp(newID, newPassword)
	} else {
		TakesLogin(newID, newPassword)
	}
}
func TakesLogin(uID, pass string) {
	if CheckCred.Contains(CheckCred.RFile(FileName), uID, pass) {
		fmt.Println("Login successfully")
	} else {
		fmt.Println("Login fail")
	}
}
func TakesSignUp(uID, pass string) {
	rd := bufio.NewReader(os.Stdin)

	fmt.Print("Enter Age: ")
	age2, _ := rd.ReadString('\n')
	age2 = strings.TrimSpace(age2)
	age, errAge := strconv.Atoi(age2)

	// Constraints applied
	if errAge != nil {
		panic(errAge)
	}
	if age >= 16 {
		CheckCred.CheckFor(FileName, uID, pass)
		fmt.Println("success")
	} else {
		fmt.Println("Unauthorized user")
	}
}

package logIn

import (
	"bufio"
	"fmt"
	"os"
	"projects1/todo"
	"projects1/util"
	"strings"
)

func LogIn() bool {
	// Enter new credentials
	fmt.Print("Enter User ID: ")
	rd := bufio.NewReader(os.Stdin)
	userID, _ := rd.ReadString('\n')
	userID = strings.TrimSpace(userID)

	fmt.Print("Enter Password: ")
	password, _ := rd.ReadString('\n')
	password = strings.TrimSpace(password)
	if UserExists(util.ReadCredFile(), userID, password) {
		fmt.Println("Login successfully")
		todo.TODO()
	} else {
		fmt.Println("Login fail")
		return false
	}

	return true
}

// UserExists returns true if user already exists and false if not
func UserExists(cred map[string]string, uID, pass string) bool {
	return cred["UserID"] == uID && cred["Password"] == pass
}

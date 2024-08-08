package logIn

import (
	"bufio"
	"fmt"
	"os"
	"projects1/admin"
	"projects1/user"
	"projects1/util"
	"strings"
)

func LogIn(appStatus string) {
	var userID, hashPassword string
	for {
		fmt.Print("Enter User ID: ")
		rd := bufio.NewReader(os.Stdin)
		userID, _ = rd.ReadString('\n')
		userID = strings.TrimSpace(userID)

		fmt.Print("Enter Password: ")
		password, _ := rd.ReadString('\n')
		password = strings.TrimSpace(password)
		hashPassword = util.HashSHA256(password)
		if appStatus == "admin" {
			if util.AdminExists(util.ReadCredFile(), userID, hashPassword) {
				fmt.Println("Login successfully for admin")
				admin.Admin()
				break
			} else {
				fmt.Println("Login fail")
			}
		} else if appStatus == "user" {
			if util.UserExists(util.ReadCredFile(), userID, hashPassword) {
				fmt.Println("Login successfully for user")
				util.DisplayToUser(userID)
				user.ManageTask(userID)
				//util.DisplayToUser(userID)
				break
			} else {
				fmt.Println("Login fail")
			}
		} else {
			fmt.Println("Login fail")
		}
	}
}

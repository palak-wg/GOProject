package user

import (
	"bufio"
	"fmt"
	"os"
	"projects1/course"
	"projects1/dailyStatus"
	"projects1/todo"
	"strings"
)

func ManageTask(uID string) {
	for {
		fmt.Println("Select an option:")
		fmt.Println("1. Manage todo")
		fmt.Println("2. Manage dailystatus")
		fmt.Println("3. Manage course")
		fmt.Println("4. Exit")

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
			todo.TODO(uID)
		case "2":
			dailyStatus.DS(uID)
		case "4":
			fmt.Println("Exiting...")
			return
		case "3":
			course.Course(uID)
		default:
			fmt.Println("Invalid option. Please choose again.")
		}
	}
}

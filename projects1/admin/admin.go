package admin

// import (
//
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//
// )
//
//	func Admin() {
//		fmt.Println("\n--------welcome admin-----------\n")
//		//util.DisplayToAdmin()
//		dataFile := "credentials.txt"
//		users, errAdmin := readUsersFromFile(dataFile)
//		if errAdmin != nil {
//			fmt.Printf("Error reading file: %v\n", errAdmin)
//			return
//		}
//
//		for {
//			fmt.Println("Admin Menu:")
//			fmt.Println("1. View user progress")
//			fmt.Println("2. Change user app status")
//			fmt.Println("3. Exit")
//
//			var choice int
//			fmt.Scanln(&choice)
//
//			switch choice {
//			case 1:
//				viewUserProgress(users)
//			case 2:
//				changeUserAppStatus(users, dataFile)
//			case 3:
//				return
//			default:
//				fmt.Println("Invalid choice. Please try again.")
//			}
//		}
//	}
//
// // Read users from the data file
//
//	func readUsersFromFile(filename string) (map[string]string, error) {
//		fileForAdmin, errForAdmin := os.Open(filename)
//		if errForAdmin != nil {
//			return nil, errForAdmin
//		}
//		defer fileForAdmin.Close()
//
//		users := make(map[string]string)
//		scannerAdmin := bufio.NewScanner(fileForAdmin)
//		for scannerAdmin.Scan() {
//			line := scannerAdmin.Text()
//			parts := strings.Split(line, ",")
//			if len(parts) != 3 {
//				return nil, fmt.Errorf("invalid line format: %s", line)
//			}
//			userID := parts[0]
//			userData := strings.Join(parts[1:], ",")
//			users[userID] = userData
//		}
//
//		if errForAdmin := scannerAdmin.Err(); errForAdmin != nil {
//			return nil, errForAdmin
//		}
//
//		return users, nil
//	}
//
// // View user progress
//
//	func viewUserProgress(users map[string]string) {
//		fmt.Println("User ID\tProgress")
//		for userID, data := range users {
//			parts := strings.Split(data, ",")
//			if len(parts) == 2 {
//				fmt.Printf("%s\t%s\n", userID, parts[0])
//			}
//		}
//	}
//
// // Change user app status
//
//	func changeUserAppStatus(users map[string]string, filename string) {
//		fmt.Println("Enter user ID to change status:")
//		var userID string
//		fmt.Scanln(&userID)
//
//		data, exists := users[userID]
//		if !exists {
//			fmt.Println("User not found.")
//			return
//		}
//
//		parts := strings.Split(data, ",")
//		if len(parts) != 2 {
//			fmt.Println("Invalid user data format.")
//			return
//		}
//
//		fmt.Printf("Current status of user %s: %s\n", userID, parts[1])
//		fmt.Println("Enter new status (user/admin):")
//		var newStatus string
//		fmt.Scanln(&newStatus)
//
//		if newStatus != "user" && newStatus != "admin" {
//			fmt.Println("Invalid status. Must be 'user' or 'admin'.")
//			return
//		}
//
//		// Update user data
//		users[userID] = fmt.Sprintf("%s,%s", parts[0], newStatus)
//		if err := writeUsersToFile(users, filename); err != nil {
//			fmt.Printf("Error updating file: %v\n", err)
//		}
//	}
//
// // Write users back to the file
//
//	func writeUsersToFile(users map[string]string, filename string) error {
//		file, err := os.Create(filename)
//		if err != nil {
//			return err
//		}
//		defer file.Close()
//
//		writer := bufio.NewWriter(file)
//		for userID, data := range users {
//			line := fmt.Sprintf("%s,%s\n", userID, data)
//			if _, err := writer.WriteString(line); err != nil {
//				return err
//			}
//		}
//
//		return writer.Flush()
//	}
func Admin() {}

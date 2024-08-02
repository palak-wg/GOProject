package CheckCred

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CheckFor(FName, uID, pass string) {
	// Read existing cred
	existCred := RFile(FName)

	// Check if already exist
	if Contains(existCred, uID, pass) {
		// Already
		fmt.Println("Credentials already exist.")
	} else {
		// Append
		if err := appendCred(FName, uID, pass); err != nil {
			fmt.Println("Error appending credentials:", err)
		} else {
			fmt.Println("User added")
		}
	}
}

func appendCred(FName, uID, pass string) error {
	file, err := os.OpenFile(FName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("UserID: " + uID + "\nPassword: " + pass + "\n")
	if err != nil {
		return err
	}

	return writer.Flush()
}

func RFile(FName string) map[string]string {
	file, err := os.Open(FName)
	if err != nil {
		if os.IsNotExist(err) {
			return make(map[string]string) // File does not exist, return empty map
		}
		return nil
	}
	defer file.Close()

	cred := make(map[string]string)
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, ": ", 2)
		if len(parts) == 2 {
			key := strings.TrimSpace(parts[0])
			value := strings.TrimSpace(parts[1])
			cred[key] = value
		}
	}

	return cred
}

func Contains(cred map[string]string, uID, pass string) bool {
	return cred["UserID"] == uID && cred["Password"] == pass
}

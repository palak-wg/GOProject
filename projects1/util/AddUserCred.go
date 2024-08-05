package util

import (
	"bufio"
	"os"
)

const FileName = "credentials.txt"

func AddUser(uID, pass string) error {
	file, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
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

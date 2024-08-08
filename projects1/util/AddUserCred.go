package util

import (
	"bufio"
	"os"
)

const FileName = "credentials.txt"

func AddUser(uID, pass, name, phonenum, address, desig string) error {
	file, err := os.OpenFile(FileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0600)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("\n" + uID + "ID: " + uID + "\n" + uID + "Password: " + pass + "\n" + uID + "AppStatus: User" + "\n" + uID + "Name: " + name + "\n" + uID + "Phone Number: " + phonenum + "\n" + uID + "Address: " + address + "\n" + uID + "Designation: " + desig + "\n")
	if err != nil {
		return err
	}
	return writer.Flush()
}

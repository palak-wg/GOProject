package util

import (
	"fmt"
)

// import (
//
//	"bufio"
//	"fmt"
//	"os"
//	"strings"
//
// )
//
//	func DisplayToAdmin() {
//		cred1 := readthisfile()
//		fmt.Println(cred1["ID"], cred1["Password"])
//		fmt.Println("hi admin")
//	}
func DisplayToUser(uID string) {
	cred1 := ReadCredFile()
	fmt.Println("---------Your Profile-----------")
	if cred1[uID+"ID"] == uID {
		fmt.Println("\nID: ", cred1[uID+"ID"], "\nName: ", cred1[uID+"Name"], "\nPhone Number: ", cred1[uID+"Phone Number"], "\nAddress: ", cred1[uID+"Address"], "\nDesignation: ", cred1[uID+"Designation"], "\n")
		fmt.Println("")
	}
	fmt.Println("hi user", uID)
}

//func readthisfile() map[string]string {
//	fileThis, errThis := os.Open(FileName)
//	if errThis != nil {
//		if os.IsNotExist(errThis) {
//			return make(map[string]string) // File does not exist, return empty map
//		}
//		return nil
//	}
//	defer fileThis.Close()
//
//	cred1 := make(map[string]string)
//	scannerThis := bufio.NewScanner(fileThis)
//
//	for scannerThis.Scan() {
//		line := scannerThis.Text()
//		parts := strings.SplitN(line, ": ", 2)
//		if len(parts) == 2 {
//			key := strings.TrimSpace(parts[0])
//			value := strings.TrimSpace(parts[1])
//			cred1[key] = value
//		}
//	}
//
//	return cred1
//}

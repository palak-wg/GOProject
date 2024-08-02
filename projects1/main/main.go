package main

import (
	"fmt"
	"projects1/InitialLog"
)

func main() {
	//file, err1 := os.Create("empty.txt")
	//
	//defer file.Close()
	//
	//if err1 != nil {
	//	log.Fatal(err1)
	//}

	fmt.Println("file created")
	var num int
	fmt.Println("Enter your choice: ")
	fmt.Println("1. SIGNUP")
	fmt.Println("2. LOGIN")
	_, errScan := fmt.Scan(&num)
	if errScan != nil {
		fmt.Println("Error:", errScan)
		return
	}

	// Use a switch statement to handle different cases
	switch num {
	case 1:
		fmt.Println("--------SIGNUP-------")
		InitialLog.Takes(1)
	case 2:
		fmt.Println("--------LOGIN--------")
		InitialLog.Takes(2)
	default:
		fmt.Println("Enter valid input")
	}
}

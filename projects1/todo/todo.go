package todo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func TODO() {

	for TodoListManipulationBool := true; TodoListManipulationBool; {

		fmt.Println("----------TODO OPTIONS:--------- \n1. ADD ITEM\n2. DELETE ITEM\n3. LIST ITEM\n4. EXIT")
		var todochoice int
		fmt.Print("Enter your choice: ")
		_, errScan := fmt.Scan(&todochoice)
		if errScan != nil {
			fmt.Println("Error:", errScan)
			return
		}
		switch todochoice {
		case 1:
			Add()
		case 2:
			Delete()
		case 3:
			List()
		case 4:
			TodoListManipulationBool = false
		default:
			fmt.Println("Invalid choice")
		}
	}

}

type TDItem struct {
	Task string
}

var todoList = []TDItem{}

func Add() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("\nEnter a new task: ")
	newTask, _ := reader.ReadString('\n')
	newTask = strings.TrimSpace(newTask)

	todoList = append(todoList, TDItem{Task: newTask})
	fmt.Println("Task added.")
}

func List() {
	if len(todoList) == 0 {
		fmt.Println("No tasks in the to-do list.")
		return
	}
	fmt.Println("\nTo-Do List:")
	for i, item := range todoList {
		fmt.Printf("%d. %s\n", i+1, item.Task)
	}
}

func Delete() {
	List()
	fmt.Print("Enter task number to delete: ")
	var listnum int
	_, err := fmt.Scan(&listnum)
	if err != nil || listnum < 1 || listnum > len(todoList) {
		fmt.Println("Invalid task number.")
		return
	}
	todoList = append(todoList[:listnum-1], todoList[listnum:]...)
	fmt.Println("Task deleted.")
}

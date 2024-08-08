package dailyStatus

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Description string
}

func DS(uID string) {
	fmt.Println("------------Welcome To DailyStatus CLI------------")
	var taskFileName = uID + "daily.txt"
	fmt.Println("------------Your DailyStatus------------")
	listDailyStatus(taskFileName)
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Insert DailyStatus")
		fmt.Println("2. View DailyStatus")
		fmt.Println("3. Exit\n")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			insertDailyStatus(taskFileName)
		case 2:
			listDailyStatus(taskFileName)
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}
}

func insertDailyStatus(filename string) {
	fmt.Print("Enter DailyStatus: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	todo := scanner.Text()

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	if _, err := file.WriteString(todo + "\n"); err != nil {
		fmt.Println("Error writing to file:", err)
	}
	fmt.Println("DailyStatus added.")
}

func listDailyStatus(filename string) {
	tasks, err := readTasks(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No DailyStatus found.")
		return
	}

	fmt.Println("DailyStatus list:")
	for i, task := range tasks {

		fmt.Printf("%d: %s \n", i+1, task.Description)
	}
}

func readTasks(filename string) ([]Task, error) {
	file, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var tasks []Task
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, "\n")
		description := parts[0]
		tasks = append(tasks, Task{Description: description})
	}

	return tasks, nil
}

func writeTasks(filename string, tasks []Task) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	for _, task := range tasks {
		_, err := file.WriteString(fmt.Sprintf("%s \n", task.Description))
		if err != nil {
			return err
		}
	}

	return nil
}

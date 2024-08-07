package todo

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Task represents a to-do item with its completion status
type Task struct {
	Description string
	Completed   bool
}

func TODO(uID string) {
	fmt.Println("------------Welcome To ToDo CLI------------")
	var taskFileName = uID + ".txt"
	fmt.Println("------------Your TODO List------------")
	listTodos(taskFileName)
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. Insert a new to-do")
		fmt.Println("2. Delete a to-do")
		fmt.Println("3. List all to-dos")
		fmt.Println("4. Mark a to-do as completed")
		fmt.Println("5. Unmark a to-do as completed")
		fmt.Println("6. Exit")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			insertTodo(taskFileName, "")
		case 2:
			deleteTodo(taskFileName, 0)
		case 3:
			listTodos(taskFileName)
		case 4:
			markCompleted(taskFileName, 0, true)
		case 5:
			markCompleted(taskFileName, 0, false)
		case 6:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}
}

func insertTodo(filename string, content string) {
	fmt.Print("Enter the new to-do item: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	content = scanner.Text()

	file, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	//if err != nil {
	//	fmt.Println("Error opening file:", err)
	//	return
	//}
	defer file.Close()

	if _, err := file.WriteString(content + " | not completed\n"); err != nil {
		fmt.Println("Error writing to file:", err)
	}
	fmt.Println("To-do added.")
}

func deleteTodo(filename string, index int) {
	fmt.Print("Enter the number of the to-do item to delete: ")
	fmt.Scan(&index)

	tasks, err := readTasks(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if index <= 0 || index > len(tasks) {
		fmt.Println("Invalid number.")
		return
	}

	tasks = append(tasks[:index-1], tasks[index:]...)
	err = writeTasks(filename, tasks)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	fmt.Println("To-do deleted.")
}

func listTodos(filename string) {
	tasks, err := readTasks(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No to-dos found.")
		return
	}

	fmt.Println("To-do list:")
	for i, task := range tasks {
		status := "not completed"
		if task.Completed {
			status = "completed"
		}
		fmt.Printf("%d: %s [%s]\n", i+1, task.Description, status)
	}
	progress := calculateProgress(tasks)
	fmt.Printf("Progress: %.2f%% completed\n", progress)
}

func calculateProgress(tasks []Task) float64 {
	total := len(tasks)
	if total == 0 {
		return 0
	}

	completed := 0
	for _, task := range tasks {
		if task.Completed {
			completed++
		}
	}

	return (float64(completed) / float64(total)) * 100
}

func markCompleted(filename string, num int, completed bool) {
	fmt.Print("Enter the number of the to-do item to update: ")
	fmt.Scan(&num)

	tasks, err := readTasks(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if num <= 0 || num > len(tasks) {
		fmt.Println("Invalid number.")
		return
	}

	tasks[num-1].Completed = completed
	err = writeTasks(filename, tasks)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return
	}

	status := "completed"
	if !completed {
		status = "not completed"
	}
	fmt.Printf("To-do marked as %s.\n", status)
}

func readTasks(filename string) ([]Task, error) {
	file, _ := os.ReadFile(filename)
	//if err != nil {
	//	return nil, err
	//}

	var tasks []Task
	lines := strings.Split(string(file), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}
		parts := strings.Split(line, " | ")
		if len(parts) != 2 {
			continue
		}
		description := parts[0]
		completed := strings.TrimSpace(parts[1]) == "completed"
		tasks = append(tasks, Task{Description: description, Completed: completed})
	}

	return tasks, nil
}

func writeTasks(filename string, tasks []Task) error {
	file, _ := os.Create(filename)
	//if err != nil {
	//	return err
	//}
	defer file.Close()

	for _, task := range tasks {
		status := "not completed"
		if task.Completed {
			status = "completed"
		}
		_, err := file.WriteString(fmt.Sprintf("%s | %s\n", task.Description, status))
		if err != nil {
			return err
		}
	}

	return nil
}

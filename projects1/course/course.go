package course

import (
	"fmt"
	"os"
	"strings"
)

// Task represents a to-do item with its completion status
type Task struct {
	Description string
	Completed   bool
}

func Course(uID string) {
	fmt.Println("------------Welcome To Course CLI------------")
	var taskFileName = uID + "course.txt"
	insertCourses(taskFileName)
	fmt.Println("------------Your Course List------------")
	listCourses(taskFileName)
	for {
		fmt.Println("\nChoose an option:")
		fmt.Println("1. List all courses")
		fmt.Println("2. Mark a course as completed")
		fmt.Println("3. Exit\n")

		var choice int
		fmt.Scan(&choice)
		switch choice {
		case 1:
			listCourses(taskFileName)
		case 2:
			markCompleted(taskFileName, true)
		case 3:
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid choice. Please choose again.")
		}
	}
}

func insertCourses(filename string) {

	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	for i := 1; i < 6; i++ {
		if _, err := file.WriteString("mod " + string(rune(i)) + " | not completed\n"); err != nil {
			fmt.Println("Error writing to file:", err)
		}
	}
}

func listCourses(filename string) {
	tasks, err := readTasks(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	if len(tasks) == 0 {
		fmt.Println("No course found.")
		return
	}

	fmt.Println("course list:")
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

func markCompleted(filename string, completed bool) {
	fmt.Print("Enter the number of the to-do item to update: ")
	var num int
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
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
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

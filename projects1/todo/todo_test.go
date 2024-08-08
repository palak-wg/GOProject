package todo

import (
	"fmt"
	"os"
	"strings"
	"testing"
)

// Helper function to write tasks to the file
func writeTasksToFile(filename string, tasks []Task) error {
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
		if _, err := file.WriteString(task.Description + " | " + status + "\n"); err != nil {
			return err
		}
	}
	return nil
}

// Helper function to read tasks from a file
func readTasksFromFile(filename string) ([]Task, error) {
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

// Table-driven test for insertTodo
func TestInsertTodo(t *testing.T) {
	tests := []struct {
		initialTasks []Task
		newTask      Task
		expected     []Task
	}{
		{
			initialTasks: []Task{},
			newTask:      Task{Description: "New Task", Completed: false},
			expected:     []Task{{Completed: false}},
		},
		{
			initialTasks: []Task{{Description: "Existing Task", Completed: false}},
			newTask:      Task{Description: "Another Task", Completed: false},
			expected:     []Task{{Description: "Existing Task", Completed: false}, {Completed: false}},
		},
		{
			initialTasks: []Task{{Description: "Existing Task", Completed: false}, {Description: "Another Task", Completed: false}},
			newTask:      Task{Description: "Another2 Task", Completed: false},
			expected:     []Task{{Description: "Existing Task", Completed: false}, {Description: "Another Task", Completed: false}, {Completed: false}},
		},
	}

	for _, tt := range tests {
		// Create a temporary file
		file, err := os.CreateTemp("", "test_insert_todo_*.txt")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(file.Name()) // Clean up the file after test

		// Write initial tasks to the file
		if err := writeTasksToFile(file.Name(), tt.initialTasks); err != nil {
			t.Fatalf("Failed to write initial tasks: %v", err)
		}

		// Insert new task
		insertTodo(file.Name(), tt.newTask.Description)
		fmt.Println(tt.newTask.Description)

		// Read the tasks and verify
		tasks, err := readTasksFromFile(file.Name())
		if err != nil {
			t.Fatalf("Failed to read tasks: %v", err)
		}
		fmt.Println(tasks)
		fmt.Println("")

		if !tasksEqual(tasks, tt.expected) {
			t.Errorf("Expected tasks %v, but got %v", tt.expected, tasks)
		}
	}
}

// Table-driven test for deleteTodo
func TestDeleteTodo(t *testing.T) {
	tests := []struct {
		initialTasks []Task
		deleteIndex  int
		expected     []Task
	}{
		{
			initialTasks: []Task{{Description: "Task 1", Completed: false}, {Description: "Task 2", Completed: false}},
			deleteIndex:  1,
			expected:     []Task{{Description: "Task 2", Completed: false}},
		},
		{
			initialTasks: []Task{{Description: "Task 1", Completed: false}},
			deleteIndex:  1,
			expected:     []Task{},
		},
		{
			initialTasks: []Task{{Description: "Task 2", Completed: false}},
			deleteIndex:  -1,
			expected:     []Task{{Description: "Task 2", Completed: false}},
		},
	}

	for _, tt := range tests {
		// Create a temporary file
		file, err := os.CreateTemp("", "test_delete_todo_*.txt")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(file.Name()) // Clean up the file after test

		// Write initial tasks to the file
		if err := writeTasksToFile(file.Name(), tt.initialTasks); err != nil {
			t.Fatalf("Failed to write initial tasks: %v", err)
		}

		// Delete the task
		deleteTodo(file.Name(), tt.deleteIndex)

		// Read the tasks and verify
		tasks, err := readTasksFromFile(file.Name())
		if err != nil {
			t.Fatalf("Failed to read tasks: %v", err)
		}

		if !tasksEqual(tasks, tt.expected) {
			t.Errorf("Expected tasks %v, but got %v", tt.expected, tasks)
		}
	}
}

// Table-driven test for markCompleted
func TestMarkCompleted(t *testing.T) {
	tests := []struct {
		initialTasks []Task
		markIndex    int
		completed    bool
		expected     []Task
	}{
		{
			initialTasks: []Task{{Description: "Task 1", Completed: false}},
			markIndex:    1,
			completed:    true,
			expected:     []Task{{Description: "Task 1", Completed: true}},
		},
		{
			initialTasks: []Task{{Description: "Task 1", Completed: false}},
			markIndex:    1,
			completed:    false,
			expected:     []Task{{Description: "Task 1", Completed: false}},
		},
		{
			initialTasks: []Task{{Description: "Task 1", Completed: false}},
			markIndex:    -1,
			completed:    false,
			expected:     []Task{{Description: "Task 1", Completed: false}},
		},
	}

	for _, tt := range tests {
		// Create a temporary file
		file, err := os.CreateTemp("", "test_mark_completed_*.txt")
		if err != nil {
			t.Fatalf("Failed to create temporary file: %v", err)
		}
		defer os.Remove(file.Name()) // Clean up the file after test

		// Write initial tasks to the file
		if err := writeTasksToFile(file.Name(), tt.initialTasks); err != nil {
			t.Fatalf("Failed to write initial tasks: %v", err)
		}

		// Mark task as completed
		markCompleted(file.Name(), tt.markIndex, tt.completed)

		// Read the tasks and verify
		tasks, err := readTasksFromFile(file.Name())
		if err != nil {
			t.Fatalf("Failed to read tasks: %v", err)
		}

		if !tasksEqual(tasks, tt.expected) {
			t.Errorf("Expected tasks %v, but got %v", tt.expected, tasks)
		}
	}
}

// Helper function to compare tasks slices
func tasksEqual(a, b []Task) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

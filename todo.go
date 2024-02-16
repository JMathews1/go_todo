package main

import (
	"fmt"
	"os"
	"encoding/gob"
)

type Task struct {
	Description string
	Done        bool
}

var tasks []Task // A slice to hold tasks

func saveTasks() error {
	file, err := os.Create("tasks.gob")
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := gob.NewEncoder(file)
	err = encoder.Encode(tasks)
	if err != nil {
		return err
	}

	return nil
}


func loadTasks() error {
	file, err := os.Open("tasks.gob")
	if err != nil {
		if os.IsNotExist(err) {
			return nil // No tasks file yet, not an error
		}
		return err
	}
	defer file.Close()

	decoder := gob.NewDecoder(file)
	err = decoder.Decode(&tasks)
	if err != nil {
		return err
	}

	return nil
}



// AddTask adds a new task with the given description
func AddTask(description string) {
	fmt.Println("addtask")
	tasks = append(tasks, Task{Description: description, Done: false})
	err := saveTasks()
	if err != nil {
		fmt.Println("Failed to save task:", err)
		return
	}
	fmt.Println("Added task:", description)
}

// ListTasks prints all tasks to the terminal
func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show!")
		return
	}
	for i, task := range tasks {
		doneStatus := "Not Done"
		if task.Done {
			doneStatus = "Done"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Description, doneStatus)
	}
}

func main() {

	err := loadTasks()
	if err != nil {
		fmt.Println("Failed to load tasks:", err)
		return
	}


	if len(os.Args) < 2 {
		fmt.Println("Usage: todo <command> [arguments]")
		fmt.Println("Commands: add, list")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: todo add <task description>")
			return
		}
		AddTask(os.Args[2])
	case "list":
		ListTasks()
	default:
		fmt.Println("Unknown command:", os.Args[1])
		fmt.Println("Available commands: add, list")
	}
}

package cli

import (
	"ToDo-CLI/model"
	"ToDo-CLI/service"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tasks []model.Task

func Run() {
	var err error
	tasks, err = service.Load()

	if err != nil {
		fmt.Println("Error loading tasks:", err)
	}

	scanner := bufio.NewScanner(os.Stdin)

	for {
		showMenu()
		if !scanner.Scan() {
			break
		}
		choice := strings.TrimSpace(scanner.Text())

		switch choice {
		case "1":
			addTask(scanner)
		case "2":
			listTasks()
		case "3":
			markTaskDone(scanner)
		case "4":
			deleteTask(scanner)
		case "5":
			err := service.Save(tasks)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			}
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Please Try again.")
		}
	}
}

func showMenu() {
	fmt.Println("\n--- TODO List ---")
	fmt.Println("1. Add Task")
	fmt.Println("2. List Tasks")
	fmt.Println("3. Mark Task as Done")
	fmt.Println("4. Delete Task")
	fmt.Println("5. Quit")
	fmt.Print("Choose an option: ")
}

func addTask(scanner *bufio.Scanner) {
	fmt.Print("Enter task description: ")

	if scanner.Scan() {
		description := strings.TrimSpace(scanner.Text())
		if description == "" {
			fmt.Println("Task cannot be empty.")

			return
		}
		tasks = append(tasks, model.New(description))
		service.Save(tasks)

		fmt.Println("Task added!")
	}
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks yet!")

		return
	}
	for i, t := range tasks {
		status := " "
		if t.Completed {
			status = "✔"
		}
		fmt.Printf("%d. [%s] %s\n", i+1, status, t.Description)
	}
}

func markTaskDone(scanner *bufio.Scanner) {
	listTasks()

	fmt.Print("Enter task number to mark as done: ")

	if scanner.Scan() {
		index, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil || index < 1 || index > len(tasks) {
			fmt.Println("Invalid input.")

			return
		}
		tasks[index-1].MarkDone()
		service.Save(tasks)

		fmt.Println("Task marked as done!")
	}
}

func deleteTask(scanner *bufio.Scanner) {
	listTasks()

	fmt.Print("Enter task number to delete: ")

	if scanner.Scan() {
		index, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		if err != nil || index < 1 || index > len(tasks) {
			fmt.Println("Invalid input.")

			return
		}
		tasks = append(tasks[:index-1], tasks[index:]...)
		service.Save(tasks)

		fmt.Println("Task deleted!")
	}
}

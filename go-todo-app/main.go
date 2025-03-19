package main 

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var tasks []string

func main() {
	reader := bufio.NewReader(os.Stdin) //buffered reader for user input

	for {
		// display menu options
		fmt.Println ("\nTo-Do List App")
		fmt.Println ("1. Add Task")
		fmt.Println ("2. View Tasks")
		fmt.Println ("3. Remove Task")
		fmt.Println ("4. Exit")
		fmt.Print ("Choose an option: ")

		//read user input
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		//handle user selection
		switch input {
		case "1":
			addTask(reader)
		case "2":
			viewTasks()
		case "3":
			removeTask(reader)
		case "4":
			fmt.Println("Exiting...") // exit the program
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

//function to add task
func addTask(reader *bufio.Reader) {
	fmt.Print("Enter task: ")
	task, _ := reader.ReadString('\n') //read task from user input
	task = strings.TrimSpace(task) // remove any trailing whitespace
	tasks = append(tasks, task) //add task
	fmt.Println("Task added successfully!")
}

// function to view all tasks
func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	fmt.Println("\nYour Tasks:")
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task) //prints task with numberings
	}
}

//func to remove task
func removeTask(reader *bufio.Reader) {
	viewTasks() //show task before removing
	if len(tasks) == 0 {
		return
	}
	fmt.Print("Enter task number to remove: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	index, err := strconv.Atoi(input) //convert user input to integer
	if err != nil|| index < 1 || index > len(tasks){
		fmt.Println("Invalid input. Please enter a valid task number.")
		return
	}
	//remove task
	tasks = append(tasks[:index-1], tasks[index:]...)
	fmt.Println("Task removed successfully!")
}

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var todolist []string

func savetofile() {
	file, err := os.Create("todolist.txt")
	if err != nil {
		fmt.Print("Error creating file: ", err)
		return
	}
	defer file.Close()
	for _, task := range todolist {
		_, err = file.WriteString(task + "\n")
		if err != nil {
			fmt.Print("Error writing to file:", err)
			return
		}
	}

	fmt.Print("Succesfully wrote to the file")
}

func loadfromfile() {
	file, err := os.Open("todolist.txt")
	if err != nil {
		if os.IsNotExist(err) {
			return
		}
		fmt.Println("Error opening file", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		todolist = append(todolist, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file: ", err)
	}

}

func add(item string) {
	todolist = append(todolist, item)
	fmt.Print("Added Successfully\n")
}

func edittask() {
	if len(todolist) == 0 {
		fmt.Println("No tasks to edit.")
		return
	}

	fmt.Print("Enter task number to edit: ")
	var index int
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	index, err := fmt.Sscanf(input, "%d", &index)
	if err != nil || index != 1 || index < 0 || index >= len(todolist) {
		fmt.Println("Invalid task number.")
		return
	}

	fmt.Print("Enter the new task: ")
	newTask, _ := reader.ReadString('\n')
	newTask = strings.TrimSpace(newTask)

	todolist[index] = newTask
	fmt.Println("Task edited successfully.")
}

func showall() {
	for i, item := range todolist {
		fmt.Println(i, ": ", item)
	}
}

func delete(index int) {
	if index < 0 || index >= len(todolist) {
		fmt.Println("Invalid task number.")
		return
	}
	todolist = append(todolist[:index], todolist[index+1:]...)
	fmt.Println("Task delete successfully.")
}

func main() {
	fmt.Println("To Do List")
	loadfromfile()

	reader := bufio.NewReader(os.Stdin)

	for {

		fmt.Printf("1.Add\n2.Showlist\n3.Edit\n4.Delete\n5.Exit\n")
		fmt.Print("Enter Your Option: ")

		option, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading option:", err)
			continue
		}
		option = strings.TrimSpace(option)
		var choice int
		_, err = fmt.Sscanf(option, "%d", &choice)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			continue
		}

		switch choice {
		case 1:
			{
				fmt.Print("Enter your task: \n")
				item, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error: ", err)
					continue
				}
				item = strings.TrimSpace(item)
				add(item)
			}
		case 2:
			{
				showall()
			}
		case 3:
			{
				edittask()
			}
		case 4:
			{
				if len(todolist) == 0 {
					fmt.Println("No tasks to delete.")
					continue
				}
				fmt.Print("Enter task number to delete: ")
				indexStr, _ := reader.ReadString('\n')
				var index int
				_, err := fmt.Sscanf(indexStr, "%d", &index)
				if err != nil || index < 0 || index >= len(todolist) {
					fmt.Println("Invalid input.")
					continue
				}
				delete(index)
			}
		case 5:
			{
				savetofile()
				fmt.Println("Exiting... Goodbye!")
				os.Exit(0)
			}

		default:
			{
				savetofile()
				fmt.Println("Exiting...")
				os.Exit(1)
			}
		}
	}
}

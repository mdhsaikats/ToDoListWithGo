package main

import (
	"bufio"
	"fmt"
	"os"
)

var todolist []string

func add(item string) {
	todolist = append(todolist, item)
	fmt.Print("Added Successfully\n")
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
	var option int
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("1.Add\n2.Showlist\n3.Delete\n4.Exit\n")
		fmt.Printf("Enter You Option: ")
		_, err := fmt.Scanf("%d", &option)
		if err != nil {
			fmt.Println("Invalid input, please enter a number.")
			reader.ReadString('\n') // clear the input buffer
			continue
		}
		switch option {
		case 1:
			{
				fmt.Print("Enter your task: ")
				item, err := reader.ReadString('\n')
				if err != nil {
					fmt.Println("Error: ", err)
					continue
				}
				item = item[:len(item)-1]
				add(item)
			}
		case 2:
			{
				showall()
			}
		case 3:
			{
				if len(todolist) == 0 {
					fmt.Println("No tasks to delete.")
					continue
				}
				fmt.Print("Enter task number to delete: ")
				var index int
				_, err := fmt.Scanf("%d", &index)
				if err != nil {
					fmt.Println("Invalid input.")
					reader.ReadString('\n')
					continue
				}
				delete(index)

			}
		default:
			{
				fmt.Println("Exiting...")
				os.Exit(1)
			}
		}
	}

}

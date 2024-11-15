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
func main() {

	fmt.Println("To Do List")
	var option int
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Printf("1.Add\n2.Showlist\n3.Delete\n4.Exit\n")
		fmt.Printf("Enter You Option: ")
		fmt.Scanf("%d", &option)

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
		default:
			{
				fmt.Println("Exiting...")
				os.Exit(1)
			}
		}
	}

}

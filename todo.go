package main

import (
	"fmt"
)

func add(todolist *[]string) {

	var item string
	fmt.Print("Enter your task: ")
	fmt.Scanln(&item)
	*todolist = append(*todolist, item)
	fmt.Print("Add Successfully\n")

}
func delete() {

}
func showall() {

}
func main() {
	var todolist []string
	fmt.Println("To Do List")
	var option int

	for {
		fmt.Printf("1.Add\n2.Showlist\n3.Delete\n4.Exit\n")
		fmt.Printf("Enter You Option: ")
		fmt.Scanf("%d", &option)

		switch option {
		case 1:
			{
				add(&todolist)
			}
		case 2:
			{
				showall()
			}
		case 3:
			{
				delete()
			}
		case 4:
			{
				fmt.Print("Exiting....")
				return
			}
		default:
			{
				fmt.Printf("Error!")
			}
		}
	}

}

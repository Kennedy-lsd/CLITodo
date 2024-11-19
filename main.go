package main

import "fmt"

type Task struct {
	Id    int
	Title string
	Done  bool
}

var tasks []Task
var storage TaskStorage

func main() {
	storage = &FileStorage{FilePath: "tasks.json"}

	var err error
	tasks, err = storage.Load()
	if err != nil {
		fmt.Println("Error loading tasks:", err)
		return
	}
	var input string
	var title string
	var id int

	for {
		fmt.Println()
		fmt.Println("Choose an action (1 for add, 2 for show, 3 for show one, 4 for delete, 5 for update task)")
		fmt.Scanln(&input)
		if input == "ex" {
			err := storage.Save(tasks)
			if err != nil {
				fmt.Println("Error saving tasks:", err)
			} else {
				fmt.Println("Tasks saved successfully!")
			}
			return
		}
		switch input {
		case "1":
			fmt.Println("Add a title:")
			fmt.Scanln(&title)
			AddTask(title)
		case "2":
			fmt.Println("Your tasks:")
			ShowTask()
		case "3":
			fmt.Println("Show by id:")
			fmt.Scanln(&id)
			ShowOneTask(id)
		case "4":
			fmt.Println("Delete by id")
			fmt.Scanln(&id)
			DeleteTask(id)
		case "5":
			fmt.Println("Update by id")
			fmt.Scanln(&id)
			UpdateTask(id)
		default:
			fmt.Println("command not found")
		}
	}

}

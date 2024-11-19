package main

import "fmt"

func AddTask(title string) {
	id := len(tasks) + 1
	task := &Task{
		Id:    id,
		Title: title,
		Done:  false,
	}
	tasks = append(tasks, *task)
	fmt.Println("Task added successfully")
}

func ShowTask() {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for _, task := range tasks {
		status := "Not completed"
		if task.Done {
			status = "Completed"
		}
		fmt.Printf("ID: %d, Title: %s, Status: %s\n", task.Id, task.Title, status)
	}
}

func ShowOneTask(id int) {
	for _, task := range tasks {
		if task.Id == id {
			status := "Not completed"
			if task.Done {
				status = "Completed"
			}
			fmt.Printf("ID: %d, Title: %s, Status: %s\n", task.Id, task.Title, status)
		} else {
			fmt.Println("Not found")
		}
	}

}

func DeleteTask(id int) {
	for i, task := range tasks {
		if task.Id == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			fmt.Println("Task deleted successfully!")
			return
		}
	}
	fmt.Println("Task not found")
}

func UpdateTask(id int) {
	for i, task := range tasks {
		if task.Id == id {
			tasks[i].Done = !tasks[i].Done
			fmt.Println("Task status updated.")
			return
		}
	}
	fmt.Println("Task not found")
}

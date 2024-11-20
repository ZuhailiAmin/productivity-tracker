package main

import (
    "fmt"
    "os"
)

func main() {
	db := initDB()
    defer db.Close() // Ensure the database connection is closed when the program ends

    if len(os.Args) < 2 {
        fmt.Println("Usage: go-productivity-tracker <command> [arguments]")
        return
    }

    switch os.Args[1] {
    case "add":
        addTask(os.Args[2:])
    case "list":
        listTasks()
    case "delete":
        deleteTask(os.Args[2:])
    default:
        fmt.Println("Unknown command")
    }
}

func addTask(args []string) {
    if len(args) < 1 {
		fmt.Println("Usage: go-productivity-tracker add <task-title>")
		return
	}

	title := args[0]
	db := initDB()
	defer db.Close()

	_, err := db.Exec("INSERT INTO tasks (title, done) VALUES (?, ?)", title, false)
	if err != nil {
		fmt.Printf("Failed to add task: %v\n", err)
		return
	}

	fmt.Println("Task added:", title)
}

func listTasks() {
    db := initDB()
	defer db.Close()

	rows, err := db.Query("SELECT id, title, done FROM tasks")
	if err != nil {
		fmt.Printf("Failed to list tasks: %v\n", err)
		return
	}
	defer rows.Close()

	fmt.Println("Tasks:")
	for rows.Next() {
		var id int
		var title string
		var done bool
		if err := rows.Scan(&id, &title, &done); err != nil {
			fmt.Printf("Failed to read task: %v\n", err)
			return
		}
		status := "Pending"
		if done {
			status = "Done"
		}
		fmt.Printf("[%d] %s - %s\n", id, title, status)
	}
}

func deleteTask(args []string) {
	if len(args) < 1 {
		fmt.Println("Usage: go-productivity-tracker delete <task-id>")
		return
	}

	id := args[0]
	db := initDB()
	defer db.Close()

	_, err := db.Exec("DELETE FROM tasks WHERE id = ?", id)
	if err != nil {
		fmt.Printf("Failed to delete task: %v\n", err)
		return
	}

	fmt.Println("Task deleted:", id)
}

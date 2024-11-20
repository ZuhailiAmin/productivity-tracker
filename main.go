package main

import (
    "fmt"
    "os"
)

func main() {
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
    fmt.Println("Adding task:", args)
}

func listTasks() {
    fmt.Println("Listing tasks...")
}

func deleteTask(args []string) {
    fmt.Println("Deleting task:", args)
}

package cmd

import (
	"cli-task-manager/storage"
	"cli-task-manager/task"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// CLI handles command-line interface
type CLI struct {
    taskManager *task.TaskManager
    storage     *storage.FileStorage
}

// NewCLI creates new CLI
func NewCLI() *CLI {
    return &CLI{
        taskManager: task.NewTaskManager(),
        storage:     storage.NewFileStorage(""),
    }
}

// Run executes CLI commands
func (c *CLI) Run() {
    // Load existing tasks
    if err := c.loadTasks(); err != nil {
        fmt.Printf("Warning: %v\n", err)
    }
    
    // Parse command-line arguments
    if len(os.Args) < 2 {
        c.showHelp()
        return
    }
    
    command := os.Args[1]
    
    switch command {
    case "add":
        c.handleAdd()
    case "list":
        c.handleList()
    case "complete":
        c.handleComplete()
    case "delete":
        c.handleDelete()
    case "help":
        c.showHelp()
    default:
        fmt.Printf("Unknown command: %s\n", command)
        c.showHelp()
    }
}

func (c *CLI) handleAdd() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: task add <title> [description]")
        return
    }
    
    title := os.Args[2]
    description := ""
    
    if len(os.Args) > 3 {
        description = strings.Join(os.Args[3:], " ")
    }
    
    task := c.taskManager.Add(title, description)
    
    if err := c.saveTasks(); err != nil {
        fmt.Printf("Error saving tasks: %v\n", err)
        return
    }
    
    fmt.Printf("✅ Task added: %d - %s\n", task.ID, task.Title)
}

func (c *CLI) handleList() {
    tasks := c.taskManager.List()
    
    if len(tasks) == 0 {
        fmt.Println("No tasks found.")
        return
    }
    
    fmt.Println("\n📋 Your Tasks:")
    fmt.Println(strings.Repeat("-", 50))
    
    for _, t := range tasks {
        status := "❌"
        if t.Completed {
            status = "✅"
        }
        
        fmt.Printf("%s [%d] %s\n", status, t.ID, t.Title)
        
        if t.Description != "" {
            fmt.Printf("    📝 %s\n", t.Description)
        }
        
        fmt.Printf("    📅 Created: %s\n", 
            t.CreatedAt.Format("2006-01-02 15:04"))
        
        if t.Completed && t.CompletedAt != nil {
            fmt.Printf("    ✅ Completed: %s\n", 
                t.CompletedAt.Format("2006-01-02 15:04"))
        }
        
        fmt.Println()
    }
}

func (c *CLI) handleComplete() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: task complete <task_id>")
        return
    }
    
    idStr := os.Args[2]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Printf("Invalid task ID: %s\n", idStr)
        return
    }
    
    err = c.taskManager.Complete(id)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    if err := c.saveTasks(); err != nil {
        fmt.Printf("Error saving tasks: %v\n", err)
        return
    }
    
    fmt.Printf("✅ Task %d marked as completed!\n", id)
}

func (c *CLI) handleDelete() {
    if len(os.Args) < 3 {
        fmt.Println("Usage: task delete <task_id>")
        return
    }
    
    idStr := os.Args[2]
    id, err := strconv.Atoi(idStr)
    if err != nil {
        fmt.Printf("Invalid task ID: %s\n", idStr)
        return
    }
    
    err = c.taskManager.Delete(id)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    if err := c.saveTasks(); err != nil {
        fmt.Printf("Error saving tasks: %v\n", err)
        return
    }
    
    fmt.Printf("🗑️  Task %d deleted!\n", id)
}

func (c *CLI) showHelp() {
//     fmt.Println(`
// 📋 Task CLI - Simple Task Manager

// Usage:
//   task <command> [arguments]

// Commands:
//   add <title> [description]  Add a new task
//   list                       List all tasks
//   complete <id>              Mark task as completed
//   delete <id>                Delete a task
//   help                       Show this help message

// Examples:
//   task add "Buy groceries" "Milk, bread, eggs"
//   task list
//   task complete 1
//   task delete 2
// `)
}

func (c *CLI) loadTasks() error {
    tasks, err := c.storage.Load()
    if err != nil {
        return err
    }
    
    c.taskManager.LoadTasks(tasks)
    return nil
}

func (c *CLI) saveTasks() error {
    tasks := c.taskManager.List()
    return c.storage.Save(tasks)
}
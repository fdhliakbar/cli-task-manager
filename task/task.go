package task

import (
	"fmt"
	"time"
)

// Task represents a single task
type Task struct {
    ID          int       `json:"id"`
    Title       string    `json:"title"`
    Description string    `json:"description"`
    Completed   bool      `json:"completed"`
    CreatedAt   time.Time `json:"created_at"`
    CompletedAt *time.Time `json:"completed_at,omitempty"`
}

// TaskManager manages collection of tasks
type TaskManager struct {
    tasks  []Task
    nextID int
}

// NewTaskManager creates new task manager
func NewTaskManager() *TaskManager {
    return &TaskManager{
        tasks:  make([]Task, 0),
        nextID: 1,
    }
}

// Add creates new task
func (tm *TaskManager) Add(title, description string) *Task {
    task := Task{
        ID:          tm.nextID,
        Title:       title,
        Description: description,
        Completed:   false,
        CreatedAt:   time.Now(),
    }
    
    tm.tasks = append(tm.tasks, task)
    tm.nextID++
    
    return &task
}

// List returns all tasks
func (tm *TaskManager) List() []Task {
    return tm.tasks
}

// Complete marks task as completed
func (tm *TaskManager) Complete(id int) error {
    for i := range tm.tasks {
        if tm.tasks[i].ID == id {
            if tm.tasks[i].Completed {
                return fmt.Errorf("task %d already completed", id)
            }
            
            now := time.Now()
            tm.tasks[i].Completed = true
            tm.tasks[i].CompletedAt = &now
            return nil
        }
    }
    return fmt.Errorf("task with ID %d not found", id)
}

// Delete removes task
func (tm *TaskManager) Delete(id int) error {
    for i, task := range tm.tasks {
        if task.ID == id {
            // Remove element from slice
            tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
            return nil
        }
    }
    return fmt.Errorf("task with ID %d not found", id)
}

// GetByID returns task by ID
func (tm *TaskManager) GetByID(id int) (*Task, error) {
    for _, task := range tm.tasks {
        if task.ID == id {
            return &task, nil
        }
    }
    return nil, fmt.Errorf("task with ID %d not found", id)
}

// LoadTasks loads tasks from slice (used by storage)
func (tm *TaskManager) LoadTasks(tasks []Task) {
    tm.tasks = tasks
    
    // Find next ID
    maxID := 0
    for _, task := range tasks {
        if task.ID > maxID {
            maxID = task.ID
        }
    }
    tm.nextID = maxID + 1
}
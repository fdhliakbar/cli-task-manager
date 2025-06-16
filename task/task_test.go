package task

import (
	"testing"
)

func TestTaskManager_Add(t *testing.T) {
    tm := NewTaskManager()
    
    task := tm.Add("Test Task", "Test Description")
    
    if task.ID != 1 {
        t.Errorf("Expected ID to be 1, got %d", task.ID)
    }
    
    if task.Title != "Test Task" {
        t.Errorf("Expected title 'Test Task', got '%s'", task.Title)
    }
    
    if task.Completed {
        t.Error("New task should not be completed")
    }
    
    if task.CreatedAt.IsZero() {
        t.Error("CreatedAt should be set")
    }
}

func TestTaskManager_Complete(t *testing.T) {
    tm := NewTaskManager()
    
    // Add a task
    task := tm.Add("Test Task", "")
    
    // Complete the task
    err := tm.Complete(task.ID)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    
    // Check if completed
    updatedTask, err := tm.GetByID(task.ID)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    
    if !updatedTask.Completed {
        t.Error("Task should be completed")
    }
    
    if updatedTask.CompletedAt == nil {
        t.Error("CompletedAt should be set")
    }
}

func TestTaskManager_CompleteNonExistent(t *testing.T) {
    tm := NewTaskManager()
    
    err := tm.Complete(999)
    if err == nil {
        t.Error("Expected error for non-existent task")
    }
}

func TestTaskManager_Delete(t *testing.T) {
    tm := NewTaskManager()
    
    // Add tasks
    task1 := tm.Add("Task 1", "")
    task2 := tm.Add("Task 2", "")
    
    // Delete first task
    err := tm.Delete(task1.ID)
    if err != nil {
        t.Errorf("Unexpected error: %v", err)
    }
    
    // Check remaining tasks
    tasks := tm.List()
    if len(tasks) != 1 {
        t.Errorf("Expected 1 task, got %d", len(tasks))
    }
    
    if tasks[0].ID != task2.ID {
        t.Errorf("Expected remaining task ID %d, got %d", task2.ID, tasks[0].ID)
    }
}
package storage

import (
	"cli-task-manager/task"
	"encoding/json"
	"fmt"
	"os"
)

const defaultFilename = "tasks.json"

// FileStorage handles file-based storage
type FileStorage struct {
    filename string
}

// NewFileStorage creates new file storage
func NewFileStorage(filename string) *FileStorage {
    if filename == "" {
        filename = defaultFilename
    }
    return &FileStorage{filename: filename}
}

// Save saves tasks to file
func (fs *FileStorage) Save(tasks []task.Task) error {
    // Convert to JSON
    data, err := json.MarshalIndent(tasks, "", "  ")
    if err != nil {
        return fmt.Errorf("failed to marshal tasks: %w", err)
    }
    
    // Write to file
    err = os.WriteFile(fs.filename, data, 0644)
    if err != nil {
        return fmt.Errorf("failed to write file: %w", err)
    }
    
    return nil
}

// Load loads tasks from file
func (fs *FileStorage) Load() ([]task.Task, error) {
    // Check if file exists
    if _, err := os.Stat(fs.filename); os.IsNotExist(err) {
        // Return empty slice if file doesn't exist
        return []task.Task{}, nil
    }
    
    // Read file
    data, err := os.ReadFile(fs.filename)
    if err != nil {
        return nil, fmt.Errorf("failed to read file: %w", err)
    }
    
    // Parse JSON
    var tasks []task.Task
    err = json.Unmarshal(data, &tasks)
    if err != nil {
        return nil, fmt.Errorf("failed to unmarshal tasks: %w", err)
    }
    
    return tasks, nil
}

// Exists checks if storage file exists
func (fs *FileStorage) Exists() bool {
    _, err := os.Stat(fs.filename)
    return !os.IsNotExist(err)
}
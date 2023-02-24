package database

import (
	"github.com/redat00/gotodo/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database struct {
	db gorm.DB
}

func NewDB() *Database {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &Database{db: *db}
}

// Create a category inside of database
func (database *Database) CreateCategory(name string, emoji string) *models.Category {
	category := models.Category{Name: name, Emoji: emoji}
	database.db.Create(category)
	return &category
}

// Get a category based on the given ID
func (database *Database) GetCategory(id string) *models.Category {
	var category models.Category
	database.db.Where("id = ?", id).First(&category)
	return &category
}

// Delete a category based on the given ID
func (database *Database) DeleteCategory(id string) {
	category := database.GetCategory(id)
	database.db.Delete(&category)
}

// Get all categories from database
func (database *Database) GetAllCategories() *[]models.Category {
	var categories []models.Category
	database.db.Find(&categories)
	return &categories
}

// Create a task status
func (database *Database) CreateTaskStatus(name string, emoji string, isDefault bool) *models.TaskStatus {
	taskstatus := models.TaskStatus{Name: name, Emoji: emoji, Default: isDefault}
	database.db.Create(&taskstatus)
	return &taskstatus
}

// Get a task status based on ID
func (database *Database) GetTaskStatus(id string) *models.TaskStatus {
	var taskstatus models.TaskStatus
	database.db.Where("id = ?", id).First(&taskstatus)
	return &taskstatus
}

// Delete a task status based on ID
func (database *Database) DeleteStatus(id string) {
	taskstatus := database.GetTaskStatus(id)
	database.db.Delete(&taskstatus)
}

// Get all task status
func (database *Database) GetAllTaskStatus() *[]models.TaskStatus {
	var taskstatuses []models.TaskStatus
	database.db.Find(&taskstatuses)
	return &taskstatuses
}

// Create a task
func (database *Database) CreateTask(name string, description string, category models.Category, status models.TaskStatus) *models.Task {
	task := models.Task{Name: name, Description: description, Category: category, Status: status}
	database.db.Create(&task)
	return &task
}

// Get a task based on ID
func (database *Database) GetTask(id string) *models.Task {
	var task models.Task
	database.db.Where("id = ?", id).First(&task)
	return &task
}

// Delete a task based on ID
func (database *Database) DeleteTask(id string) {
	task := database.GetTask(id)
	database.db.Delete(&task)
}

// Get all tasks
func (database *Database) GetAllTasks() *[]models.Task {
	var tasks []models.Task
	database.db.Find(&tasks)
	return &tasks
}

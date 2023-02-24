package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Name  string
	Emoji string
}

type TaskStatus struct {
	gorm.Model
	Name    string
	Emoji   string
	Default bool
}

type Task struct {
	gorm.Model
	Name        string
	Description string
	CategoryID  string
	Category    Category `gorm:"foreignKey:CategoryID;references:ID"`
	StatusID    string
	Status      TaskStatus `gorm:"foreignKey:StatusID;references:ID"`
}

type GotodoConfiguration struct {
	BaseDir string `yaml:"base_dir"`
}

package models

import "time"

type Task struct {
	ID           uint      `gorm:"primaryKey;autoIncrement;not null"`
	Title        string    `gorm:"size:255;not null"`
	Description  string    `gorm:"type:text"`
	Completed    bool      `gorm:"type:bool;default:false"`
	DateTimeTask time.Time `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

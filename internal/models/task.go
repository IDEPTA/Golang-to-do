package models

import "time"

type Task struct {
	ID           int64     `gorm:"primaryKey;autoIncrement;not null"`
	Title        string    `gorm:"size:255;not null"`
	Description  string    `gorm:"type:text"`
	Completed    bool      `gorm:"type:bool;default:false"`
	UserID       int64     `gorm:"not null"`
	DateTimeTask time.Time `gorm:"not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

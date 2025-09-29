package requests

import "time"

type TaskRequest struct {
	Title        string    `json:"title" binding:"required,min=3,max=100"`
	Description  string    `json:"description" binding:"omitempty,max=500"`
	Completed    bool      `json:"completed" binding:"omitempty"`
	DateTimeTask time.Time `json:"date_time_task" binding:"required"`
	UserID       int64     `json:"user_id" binding:"required"`
}

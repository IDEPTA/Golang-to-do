package models

type User struct {
	ID         int64  `json:"id" gorm:"primaryKey;autoIncrement"`
	Name       string `json:"name" gorm:"not null;required;min=2;max=100"`
	Lastname   string `json:"lastname" gorm:"not null;required;min=2;max=100"`
	Patronymic string `json:"patronymic" gorm:"nullable;max=100"`
	Email      string `json:"email" gorm:"unique;not null;required;email;max=100"`
	Password   string `json:"-" gorm:"not null;required;min=6;max=100"`
	Age        int    `json:"age" gorm:"nullable;min=0;max=150"`
	Tasks      []Task `json:"tasks" gorm:"foreignKey:UserID"`
}

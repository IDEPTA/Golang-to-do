package requests

type RegisterRequest struct {
	Name            string `json:"name" binding:"required,min=3,max=50"`
	Lastname        string `json:"lastname" binding:"required,min=3,max=50"`
	Patronymic      string `json:"patronymic" binding:"omitempty,max=50"`
	Email           string `json:"email" binding:"required,email"`
	Age             int    `json:"age" binding:"omitempty,min=0,max=150"`
	Password        string `json:"password" binding:"required,min=6,max=100"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password"`
}

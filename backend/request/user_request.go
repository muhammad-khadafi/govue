package request

type UserRequest struct {
	Id       uint   `json:"id" db:"ID"`
	Username string `json:"username" binding:"required,min=5,max=20" db:"USERNAME"`
	Password string `json:"password" binding:"required,max=100" db:"PASSWORD"`
	Email    string `json:"email" binding:"required,email" db:"EMAIL"`
}

package entity

type Role string

const (
	AdminRole    Role = "admin"
	CustomerRole Role = "customer"
	UserRole     Role = "user"
)

type User struct {
	ID       uint   `json:"id"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     Role   `json:"role"`
}

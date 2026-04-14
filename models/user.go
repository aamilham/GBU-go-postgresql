package models
//model user yang digunakan, masih sama seperti model user pada php slim sebelumnya, yaitu name, email, dan role, serta created_at untuk kebutuhan audit trail, dan id sebagai primary key
type User struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Role      string `json:"role"`
	CreatedAt string `json:"created_at"`
}
package users

type User struct {
	Username string `json:"username" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Gender   string `json:"gender"`
}

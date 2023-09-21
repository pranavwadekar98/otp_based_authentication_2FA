package entity

type SignUpData struct {
	Phone    string `gorm:"primary_key type:varchar(20)" json:"phone" binding:"required"`
	Password string `gorm:"type:varchar(20)" json:"password" binding:"required"`
}

// type LoginData struct {
// 	Email    string `json:"email" binding:"required,email"`
// 	Password string `json:"password" binding:"required"`
// }

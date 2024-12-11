package models

type User struct {
	Id       uint64 `gorm:"primaryKey;column:id;autoIncrement;comment:id" json:"id"`
	UserID   string `json:"user_id"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

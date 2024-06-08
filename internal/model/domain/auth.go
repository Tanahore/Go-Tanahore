package domain

import "time"

type Users struct {
	ID        uint      `gorm:"type:int;primarykey" json:"id"`
	Username  string    `gorm:"type:varchar(255)" json:"username"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	Password  string    `gorm:"type:varchar(255)" json:"password"`
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
}

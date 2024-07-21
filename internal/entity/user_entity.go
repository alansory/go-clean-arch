package entity

import "time"

type User struct {
	ID                int64      `gorm:"column:id;primaryKey"`
	Fullname          string     `gorm:"column:fullname"`
	Username          string     `gorm:"column:username"`
	Email             string     `gorm:"column:email"`
	Phone             string     `gorm:"column:phone"`
	Address           string     `gorm:"column:address"`
	HashedPassword    string     `gorm:"column:hashed_password"`
	PasswordChangedAt *time.Time `gorm:"column:password_changed_at"`
	CreatedAt         *time.Time `gorm:"column:created_at"`
	UpdatedAt         *time.Time `gorm:"column:updated_at"`
}

func (a *User) TableName() string {
	return "users"
}

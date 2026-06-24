package model

type User struct {
	ID        uint   `gorm:"primaryKey;autoIncrement"`
	Name      string
	Age       uint
}
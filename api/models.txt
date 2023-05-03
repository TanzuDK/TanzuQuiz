package models

type Users_Post struct {
	email string `gorm:"not null" json:"content,omitempty"`
}

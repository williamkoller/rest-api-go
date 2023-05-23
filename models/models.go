package models

import "time"

type Fact struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Question  string    `json:"question" gorm:"text;not null;default:null`
	Answer    string    `json:"answer" gorm:"text;not null;default:null`
}

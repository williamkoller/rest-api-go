package models

import "time"

type Fact struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Question  string    `json:"question" gorm:"text"`
	Answer    string    `json:"answer" gorm:"text"`
}

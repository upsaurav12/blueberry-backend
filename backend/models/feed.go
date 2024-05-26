package models

import "time"

type Feed struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	UserID      uint      `json:"user_id" gorm:"not null"` // Foreign Key referencing User table
	TextContent string    `json:"text_content" gorm:"type:text"`
	Timestamp   time.Time `json:"timestamp" gorm:"autoCreateTime"`
	Image       string    `json:"image"`
	ProfileURL  string    `json:"profile_url"`
	User        User      `json:"user" gorm:"foreignKey:UserID"`
}

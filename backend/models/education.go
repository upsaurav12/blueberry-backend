package models

import "time"

type Education struct {
	EducationId  uint      `json:"education_id" gorm:"primary_key"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	Institution  string    `json:"institution"`
	Degree       string    `json:"degree"`
	FieldOfStudy string    `json:"fieldofstudy"`
	StartDate    time.Time `json:"startdate" gorm:"autoCreateTime"`
	EndDate      time.Time `json:"enddate" gorm:"enddate"`
	User         User      `json:"user" gorm:"foreignKey:UserID"`
}

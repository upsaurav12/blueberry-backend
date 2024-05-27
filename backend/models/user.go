package models

type User struct {
	ID           uint    `json:"id" gorm:"primary_key"`
	UserImage    string  `json:"user_image"`
	FirstName    string  `json:"first_name"`
	LastName     string  `json:"last_name"`
	Summary      string  `json:"summary"`
	Email        string  `json:"email" gorm:"unique"`
	Location     string  `json:"location"`
	Role         string  `json:"role"`
	PhoneNumber  string  `json:"phone_number"`
	ActiveStatus bool    `json:"active_status"`
	LinkedIn     string  `json:"linkedin"`
	Instagram    string  `json:"instagram"`
	Password     string  `json:"password"`
	Feeds        []Feed  `json:"feeds" gorm:"foreignKey:UserID"`
	Groups       []Group `json:"feeds" gorm:"foreignKey:UserID"`
	/*Educations   []Education `json:"educations" gorm:"foreignKey:UserId"`*/
}

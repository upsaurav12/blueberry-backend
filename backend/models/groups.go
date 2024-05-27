package models

type Group struct {
	GroupId       uint   `json:"group_id" gorm:"primary_key"`
	UserID        uint   `json:"user_id" gorm:"not null"`
	GroupName     string `json:"group_name"`
	AdminUserId   uint   `json:"admin_id"`
	IsActive      bool   `json:"isactive"`
	Description   string `json:"description"`
	GroupImageURL string `json:"groupimage"`
	User          User   `json:"user" gorm:"foreignKey:UserID"`
}

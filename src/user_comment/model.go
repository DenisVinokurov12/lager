package user_comment

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var UserLogCallbackAdd = func(string , int, int){}
var UserLogCallbackDelete = func(string , int, int){}
var UserLogCallbackEdit = func(string , int, int){}


type UserComment struct {
	Id             	int    `gorm:"primary_key" json:"id"`
	UserId 			int    `json:"user_id"`
	IssueDayId 		int    `json:"issue_day_id"`
	UserCommentId 	int    `json:"user_comment_id"`
	Comment 		string    `gorm:"type:text" json:"comment"`
	Video 			string    `json:"video"`
	Image 			string    `json:"image"`
	File 			string    `json:"file"`
}
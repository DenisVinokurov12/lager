package course_day

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var UserLogCallbackAdd = func(string , int, int){}
var UserLogCallbackDelete = func(string , int, int){}
var UserLogCallbackEdit = func(string , int, int){}

type CourseDay struct {
	Id             	int    `gorm:"primary_key" json:"id"`
	IssueDayId      int `json:"issue_day_id"`
	UserId    		int `json:"user_id"`
	ParentUserId 	string    `json:"parent_user_id"`
	Image 			string    `json:"image"`
	AddTs 			time.Time    `json:"add_ts"`
}

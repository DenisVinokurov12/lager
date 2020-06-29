package user_course_answer

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var UserLogCallbackAdd = func(string , int, int){}
var UserLogCallbackDelete = func(string , int, int){}
var UserLogCallbackEdit = func(string , int, int){}

type UserCourseAnswer struct {
	Id             	int    		`gorm:"primary_key" json:"id"`
	UserId 			int    		`json:"day_id"`
	TypeValue 		string    		`json:"type_value"`
	IssueDayId 	int    	`json:"issue_day_id"`
	Value 		string   `json:"value"`
}

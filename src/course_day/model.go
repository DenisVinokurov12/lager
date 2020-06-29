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
	Day          	int `json:"day"`
	CourseId    	int `json:"course_id"`
	Title 			string    `json:"title"`
	Description 	string    `gorm:"type:text" json:"description"`
}

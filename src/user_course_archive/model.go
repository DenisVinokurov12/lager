package user_course_archive

import (
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

var UserLogCallbackAdd = func(string , int, int){}
var UserLogCallbackDelete = func(string , int, int){}
var UserLogCallbackEdit = func(string , int, int){}

type UserCourseArchive struct {
	Id             	int    		`gorm:"primary_key" json:"id"`
	UserId 			int    		`json:"day_id"`
	Day 			int    		`json:"day"`
	CourseId 			int    		`json:"course_id"`
	IsCompleted 	bool    	`json:"video"`
	StartTs 		time.Time   `json:"add_ts"`
}
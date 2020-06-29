package issue_day

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var UserLogCallbackAdd = func(string , int, int){}
var UserLogCallbackDelete = func(string , int, int){}
var UserLogCallbackEdit = func(string , int, int){}


type IssueDay struct {
	Id             	int    `gorm:"primary_key" json:"id"`
	DayId 			int    `json:"day_id"`
	CourseId 		int    `json:"course_id"`
	Title 			string    `json:"title"`
	Description 	string    `gorm:"type:text" json:"description"`
	Video 			string    `json:"video"`
	Image 			string    `json:"image"`
	FullDescription string    `gorm:"type:text" json:"full_description"`
	Sort 			int    		`json:"sort"`
	IsCompleted 	bool    `gorm:"-" json:"is_completed"`
}

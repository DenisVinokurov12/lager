package user_course

import (
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

var UserLogCallbackAdd = func(string , int, int){}
var UserLogCallbackDelete = func(string , int, int){}
var UserLogCallbackEdit = func(string , int, int){}

type UserCourse struct {
	Id             	int    		`gorm:"primary_key" json:"id"`
	UserId 			int    		`json:"day_id"`
	IssueDayId 		int    		`json:"issue_day_id"`
	IsCompleted 	bool    	`json:"is_completed"`
	StartTs 		time.Time   `json:"add_ts"`
}
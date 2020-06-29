package userlog

import (
	"github.com/jinzhu/gorm"
	"time"
)

var DB *gorm.DB

const(
	DELETE = "delete"
	EDIT = "edit"
	ADD = "add"
)

type UserLog struct {
	Id       int       `gorm:"primary_key" json:"id"`
	Action   string    `json:"action"`
	Object   string    `json:"object"`
	ObjectId int       `json:"object_id"`
	UserId   int       `json:"user_id"`
	AddTs    time.Time `json:"add_ts"`
}

type User interface {
	GetId() int
}


func AddToLog(model string, object_id, user_id int) {
	if user_id == -1{
		return
	}
	ulog := &UserLog{
		Action : ADD,
		Object : model,
		ObjectId : object_id,
		UserId : user_id, 
	}
	Add(ulog)
}

func DeleteToLog(model string, object_id, user_id int) {
	if user_id == -1{
		return
	}
	ulog := &UserLog{
		Action : DELETE,
		Object : model,
		ObjectId : object_id,
		UserId : user_id, 
	}
	Add(ulog)
}

func EditToLog(model string, object_id, user_id int) {
	if user_id == -1{
		return
	}
	ulog := &UserLog{
		Action : EDIT,
		Object : model,
		ObjectId : object_id,
		UserId : user_id, 
	}
	Add(ulog)
}
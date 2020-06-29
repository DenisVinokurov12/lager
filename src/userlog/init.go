package userlog

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func (UserLog) TableName() string {
	return "user_log"
}

func Init(db *gorm.DB) {
	logrus.Info("Лог пользователя")
	DB = db
	db.AutoMigrate(&UserLog{})
}

func SetEnv() {
	logrus.Info("Лог пользователя")
}

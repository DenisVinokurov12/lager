package user_course

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"os"
)

func (UserCourse) TableName() string {
	return "user_course"
}

func Init(db *gorm.DB) {
	logrus.Info("Курс пользователей")
	DB = db
	db.AutoMigrate(&UserCourse{})
}

const DIR_PREVIEW = "./public/user_source/files"

func SetEnv() {
	logrus.Info("Курс пользователей")
	os.MkdirAll(DIR_PREVIEW, os.ModePerm)
}

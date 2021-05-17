package user_course_answer

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"os"
)

func (UserCourseAnswer) TableName() string {
	return "user_course_archive"
}

func Init(db *gorm.DB) {
	logrus.Info("Ответы пользователей")
	DB = db
	db.AutoMigrate(&UserCourseAnswer{})
}

const DIR_ANSWER = "./public/user/answer"

func SetEnv() {
	logrus.Info("Ответы пользователей")
	os.MkdirAll(DIR_ANSWER, os.ModePerm)
}

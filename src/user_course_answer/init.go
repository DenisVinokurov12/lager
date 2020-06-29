package user_course_answer

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func (UserCourseAnswer) TableName() string {
	return "user_course_archive"
}

func Init(db *gorm.DB) {
	logrus.Info("Ответы пользователей")
	DB = db
	db.AutoMigrate(&UserCourseAnswer{})
}

func SetEnv() {
	logrus.Info("Ответы пользователей")
}

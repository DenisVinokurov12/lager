package user_course_archive

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func (UserCourseArchive) TableName() string {
	return "user_course_archive"
}

func Init(db *gorm.DB) {
	logrus.Info("Архив курсов пользователей")
	DB = db
	db.AutoMigrate(&UserCourseArchive{})
}

func SetEnv() {
	logrus.Info("Архив курсов пользователей")
}

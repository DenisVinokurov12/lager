package course_day

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func (CourseDay) TableName() string {
	return "course_day"
}

func Init(db *gorm.DB) {
	logrus.Info("Дни курсов")
	DB = db
	db.AutoMigrate(&CourseDay{})
}

func SetEnv() {
	logrus.Info("Дни курсов")
}

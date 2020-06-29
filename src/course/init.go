package course

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"os"
)

func (Course) TableName() string {
	return "course"
}

func Init(db *gorm.DB) {
	logrus.Info("Курсы")
	DB = db
	db.AutoMigrate(&Course{})
}

const DIR_PREVIEW = "./public/course/preview"

func SetEnv() {
	logrus.Info("Курсы")
	os.MkdirAll(DIR_PREVIEW, os.ModePerm)
}

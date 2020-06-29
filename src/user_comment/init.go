package user_comment

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func (UserComment) TableName() string {
	return "user_comment"
}

func Init(db *gorm.DB) {
	logrus.Info("Комментарии")
	DB = db
	db.AutoMigrate(&UserComment{})
}

const DIR_PREVIEW = "./public/issue_day/preview"

func SetEnv() {
	logrus.Info("Комментарии")
}

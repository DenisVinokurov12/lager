package user_comment

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"os"
)

func (UserComment) TableName() string {
	return "user_comment"
}

const DIR_COMMENT = "./public/user/comment"
const DIR_NAGRADA = "./public/user/nagrada"


func Init(db *gorm.DB) {
	logrus.Info("Комментарии")
	DB = db
	db.AutoMigrate(&UserComment{})
}

func SetEnv() {
	logrus.Info("Комментарии")
	os.MkdirAll(DIR_COMMENT, os.ModePerm)
	os.MkdirAll(DIR_NAGRADA, os.ModePerm)
}

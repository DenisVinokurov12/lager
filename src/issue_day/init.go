package issue_day

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"os"
)

func (IssueDay) TableName() string {
	return "issue_day"
}

func Init(db *gorm.DB) {
	logrus.Info("Задания на день курса")
	DB = db
	db.AutoMigrate(&IssueDay{})
}

const DIR_PREVIEW = "./public/issue_day/preview"
const DIR_FILE = "./public/issue_day/file"

func SetEnv() {
	logrus.Info("Задания на день курса")
	os.MkdirAll(DIR_PREVIEW, os.ModePerm)
	os.MkdirAll(DIR_FILE, os.ModePerm)
}

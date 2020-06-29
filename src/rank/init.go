package rank

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func (Rank) TableName() string {
	return "rank"
}

func Init(db *gorm.DB) {
	logrus.Info("Ранги")
	DB = db
	db.AutoMigrate(&Rank{})
}

func SetEnv() {
	logrus.Info("Ранги")
}

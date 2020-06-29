package order

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
)

func (Order) TableName() string {
	return "order_user"
}

func Init(db *gorm.DB) {
	logrus.Info("Покупки")
	DB = db
	db.AutoMigrate(&Order{})
}

func SetEnv() {
	logrus.Info("Покупки")
}

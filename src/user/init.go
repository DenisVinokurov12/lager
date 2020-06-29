package user

import (
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"context"
	"os"
)

func (User) TableName() string {
	return "user"
}

const DIR_AVATAR = "./public/user/avatars"

func Init(db *gorm.DB) {
	logrus.Info("Пользователи")
	DB = db
	db.AutoMigrate(&User{})
}

func SetEnv() {
	logrus.Info("Пользователи")

	os.MkdirAll(DIR_AVATAR, os.ModePerm)

	if AllLength() == 0 {

		users := []*User{&User{
			Login:    "marina777",
			Password: "violet22ilalova",
			Rule: 2,
			FirstName : "Марина",
			LastName : "Илалова",
		},&User{
			Login:    "curator111",
			Password: "violet111curator",
			Rule: 1,
			FirstName : "Куратор",
			LastName : "1",
		}}

		ctx_ := context.WithValue(context.Background(), "init_user", -1)

		for i := 0; i < len(users); i++ {
			users[i].Id = i + 1
			users[i].Password = CompressPass(users[i].Password)
			Add(users[i], ctx_ )
		}
	}
}

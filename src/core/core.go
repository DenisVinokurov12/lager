package core

import (
	"fmt"
	"github.com/jinzhu/gorm"
	// "regexp"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"errors"
	"github.com/sirupsen/logrus"
	"user"
	"user_comment"
	"user_course_archive"
	"user_course_answer"
	"user_course"
	"order"
	"userlog"
	"course"
	"rank"
	"course_day"
	"issue_day"
)

type Database struct {
	Driver string `json:"driver"`
	Path   string `json:"path"`
	Login  string `json:"login"`
	Pass   string `json:"pass"`
	Host   string `json:"host"`
	Port   int    `json:"port"`
	DBName string `json:"dbname"`
	Debug  bool   `json:"debug"`
}

var CONFIG_PATH string = ""

func New(c *Database, db *gorm.DB) *gorm.DB {

	if db == nil {
	RereadConfig:

		err := errors.New("")

		db, err = gorm.Open(c.Driver, get_connect_string_databsse(c))
		db.LogMode(c.Debug)

		if err != nil && err.Error() != "" {

			logrus.Warn("Не получилось подключиться к БД")

			cr := &Database{
				Driver: c.Driver,
				Path:   c.Path,
				Login:  c.Login,
				Pass:   c.Pass,
				Host:   c.Host,
				Port:   c.Port,
				DBName: c.DBName,
				Debug:  c.Debug,
			}
			cr.DBName = ""
			db_race, err_race := sql.Open(c.Driver, get_connect_string_databsse(cr))
			if err_race != nil {
				logrus.Warn("Не получилось подключиться для создания БД")
				return db
			}
			logrus.Info("Подключились для создания БД")
			defer db_race.Close()
			_, e1 := db_race.Exec(fmt.Sprintf("CREATE DATABASE `%s` default charset=utf8", c.DBName))
			if e1 == nil {
				logrus.Info("Создали БД")
				logrus.Info("Переподключаемся")
				goto RereadConfig
			} else {
				logrus.Fatal("Не смогли создать БД: ", e1.Error(), c)
				return db
			}
			return db
		}
	}

	db.Set("gorm:table_options", "charset=utf8")

	// if c.Driver != "sqlite3"{
	// 	rows, err_v := db.DB().Query(`select version()`)
	// 	if err_v != nil {
	// 		logrus.Fatal(err_v.Error())
	// 	}
	// 	defer rows.Close()

	// 	var version_mysql string
	// 	rows.Next()
	// 	rows.Scan(&version_mysql)

	// 	re_version := regexp.MustCompile(`([0-9|.]*)-(.*)`)
	// 	s := re_version.FindAllStringSubmatch(version_mysql, -1)

	// 	if s[0][1] < "5.7" {
	// 		logrus.Fatal("Версия MySQL не 5.7 и выше")
	// 	}
	// }

	logrus.Info("Подключение модулей:")

	user.Init(db)
	userlog.Init(db)
	course.Init(db)
	course_day.Init(db)
	issue_day.Init(db)
	user_comment.Init(db)
	rank.Init(db)
	user_course.Init(db)
	user_course_archive.Init(db)
	user_course_answer.Init(db)
	order.Init(db)

	logrus.Info("Инициализация ядра:")
	user.SetEnv()
	userlog.SetEnv()
	course.SetEnv()
	course_day.SetEnv()
	issue_day.SetEnv()
	user_comment.SetEnv()
	user_course.SetEnv()
	user_course_archive.SetEnv()
	user_course_answer.SetEnv()
	rank.SetEnv()
	order.SetEnv()

	user.UserLogCallbackAdd = userlog.AddToLog
	user.UserLogCallbackDelete = userlog.DeleteToLog
	user.UserLogCallbackEdit = userlog.EditToLog

	course.UserLogCallbackAdd = userlog.AddToLog
	course.UserLogCallbackDelete = userlog.DeleteToLog
	course.UserLogCallbackEdit = userlog.EditToLog
	
	course_day.UserLogCallbackAdd = userlog.AddToLog
	course_day.UserLogCallbackDelete = userlog.DeleteToLog
	course_day.UserLogCallbackEdit = userlog.EditToLog
	
	user_course.UserLogCallbackAdd = userlog.AddToLog
	user_course.UserLogCallbackDelete = userlog.DeleteToLog
	user_course.UserLogCallbackEdit = userlog.EditToLog
	
	issue_day.UserLogCallbackAdd = userlog.AddToLog
	issue_day.UserLogCallbackDelete = userlog.DeleteToLog
	issue_day.UserLogCallbackEdit = userlog.EditToLog
	
	order.UserLogCallbackAdd = userlog.AddToLog
	order.UserLogCallbackDelete = userlog.DeleteToLog
	order.UserLogCallbackEdit = userlog.EditToLog

	return db

}

func get_connect_string_databsse(c *Database) string {
	switch c.Driver {
	case "mysql":
		return fmt.Sprintf(`%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=UTC`,
			c.Login,
			c.Pass,
			c.Host,
			c.Port,
			c.DBName)
	case "mssql":
		return fmt.Sprintf(`sqlserver://%s:%s@%s:%d?database=%s`,
			c.Login,
			c.Pass,
			c.Host,
			c.Port,
			c.DBName)
	case "sqlite3", "sqlite":
		return c.Path
	case "postgres":
		return fmt.Sprintf(`host=%s port=%d user=%s dbname=%s password=%s`,
			c.Host,
			c.Port,
			c.Login,
			c.DBName,
			c.Pass)
	}
	logrus.Fatal(c.Driver, "undefined")
	return ""
}

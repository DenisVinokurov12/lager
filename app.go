package main

import (
	"apiserver"
	"core"
	"encoding/json"
	"flag"
	"github.com/sirupsen/logrus"
	"io/ioutil"
)

type (
	Email struct {
		Server   string `json:"server"`
		Port     int    `json:"port"`
		Login    string `json:"login"`
		Name     string `json:"name"`
		Password string `json:"password"`
	}
	Config struct {
		UsedPort             int            `json:"used_port"`
		Error                string         `json:"error"`
		NetAddr              string         `json:"net_addr"` // Адрес в интернете
		SessionKey           string         `json:"session_key"`
		Email                `json:"email"` // Настройка почты
		*core.Database       `json:"database"`
		*apiserver.ApiConfig `json:"api_config"`
	}
)

func main() {

	logrus.Info("Стартует приложение")

	config_path := flag.String("config", "", "")
	testing_auth := flag.Bool("is_testing_auth", false, "")
	flag.Parse()

	logrus.Info("Разбор аргументов")

	apiserver.IS_TESTING_AUTH = *testing_auth

	///// READ CONFIG
	logrus.Info("Читаем конфиг")
	byteValue, err_config_file := ioutil.ReadFile(*config_path)

	if err_config_file != nil {
		logrus.Fatal("Read config file: ", err_config_file.Error())
		return
	}


	config := Config{}

	json.Unmarshal(byteValue, &config)
	///// READ CONFIG

	logrus.Info("Стартует ядро системы")
	_gorm := core.New(config.Database, nil)

	logrus.Info("Стартует апи сервер")

	sql_time_diff := `select timediff(now(),convert_tz(now(),@@session.time_zone,"+00:00"))`
	rows, err := _gorm.DB().Query(sql_time_diff)
	if err != nil{
		logrus.Warnf("Не смогли получить часовой пояс: %s" , err.Error())
	}else{
		defer rows.Close()
		rows.Next()
		rows.Scan(&apiserver.DiffTime)
	}

	logrus.Infof("DiffTime: %s" , apiserver.DiffTime)
	

	apiserver.New(config.ApiConfig)

	logrus.Info("Готово к работе")

	select {}
}

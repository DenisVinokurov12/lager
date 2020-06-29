package order

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var UserLogCallbackAdd = func(string , int, int){}
var UserLogCallbackDelete = func(string , int, int){}
var UserLogCallbackEdit = func(string , int, int){}


type Order struct {
	Id             	int    `gorm:"primary_key" json:"id"`
	OrderId 			string    `json:"order_id"`
	Quantity 		int    `json:"quantity"`
}
package rank

import (
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

var UserLogCallbackAdd = func(string , int, int){}
var UserLogCallbackDelete = func(string , int, int){}
var UserLogCallbackEdit = func(string , int, int){}


type Rank struct {
	Id             	int    `gorm:"primary_key" json:"id"`
	Title 			string    `json:"title"`
	Weight 			int    `json:"weight"`
}
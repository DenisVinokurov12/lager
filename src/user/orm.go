package user

import (
	"github.com/google/uuid"
	"context"
)

func (u *User) Save() {
	DB.Create(u)
}

func AllLength() int {
	records := []User{}
	count := 0
	DB.Find(&records).Count(&count)
	return count
}

func GetByOrderId(order_id string) []*User {
	record := []*User{}
	DB.Where("order_id = ?", order_id).
		Find(&record)
	return record
}

func All() []User {
	records := []User{}
	DB.Find(&records)
	return records
}

func Update(u *User, ctx context.Context) *User {
	if u.Id == 0 {
		id, _ := uuid.NewUUID()
		u.ApiToken = id.String()
	}
	DB.Save(&u)

	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackEdit("user", u.Id, user_id)

	return u
}

// Сюда передаются оригиналы и логина и пароля
func GetByLoginByPass(login, pass string) User {
	record := User{}
	DB.Where("login = ?", login).
		Where("password = ?", CompressPass(pass)).
		First(&record)
	return record
}

func GetByPersonalArea(personal_area_id int) []User {
	record := []User{}
	DB.Where("personal_area_id = ?", personal_area_id).
		Find(&record)
	return record
}

func GetById(id int) *User {
	record := &User{}
	DB.Where("id = ?", id).
		First(&record)
	return record
}

func GetByApiToken(api_token string) User {
	record := User{}
	DB.Where("api_token = ?", api_token).
		First(&record)
	return record
}

func Add(u *User, ctx context.Context) *User {

	id, _ := uuid.NewUUID()
	u.ApiToken = id.String()
	DB.Create(u)

	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackAdd("user", u.Id, user_id)

	return u
}

func Delete(id int, ctx context.Context) *User {
	DB.Where("id = ?", id).Delete(User{})

	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackDelete("user", id, user_id)

	return GetById(id)
}


func (u *User) GetAccess() string {
	return "{}"
}
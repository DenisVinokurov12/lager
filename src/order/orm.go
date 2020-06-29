package order

import(
	"context"
)


func GetById(id int) *Order {
	record := &Order{}
	DB.Where("id = ?", id).
		First(&record)
	return record
}

func GetByOrderId(order_id string) *Order {
	record := &Order{}
	DB.Where("order_id = ?", order_id).
		First(&record)
	return record
}

func Delete(id int, ctx context.Context) *Order {
	DB.Where("id = ?", id).Delete(Order{})
	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackDelete("order", id, user_id)

	return GetById(id)
}

func Update(u *Order, ctx context.Context) *Order {
	DB.Save(&u)
	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackEdit("order", u.Id, user_id)
	return u
}

func Add(u *Order, ctx context.Context) *Order {

	DB.Create(u)

	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackAdd("order", u.Id, user_id)

	return u
}

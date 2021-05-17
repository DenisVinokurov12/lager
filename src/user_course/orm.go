package user_course

import(
	"context"
)


func GetById(id int) *UserCourse {
	record := &UserCourse{}
	DB.Where("id = ?", id).
		First(&record)
	return record
}


func GetByIdByUserId(issue_day_id, user_id int) *UserCourse {
	record := &UserCourse{}
	DB.Where("issue_day_id = ?", issue_day_id).
		Where("user_id = ?", user_id).
		First(&record)
	return record
}

func GetByUserIdNotCompleted(user_id int) *UserCourse {
	record := &UserCourse{}
	DB.Where("is_completed = ?", false).
		Where("user_id = ?", user_id).
		First(&record)
	return record
}

func Delete(id int, ctx context.Context) *UserCourse {
	DB.Where("id = ?", id).Delete(UserCourse{})
	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackDelete("user_course", id, user_id)

	return GetById(id)
}

func Update(u *UserCourse, ctx context.Context) *UserCourse {
	DB.Save(&u)
	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackEdit("user_course", u.Id, user_id)
	return u
}

func Add(u *UserCourse, ctx context.Context) *UserCourse {

	DB.Create(u)

	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackAdd("user_course", u.Id, user_id)

	return u
}

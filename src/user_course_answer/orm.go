package user_course_answer

import(
	"context"
)


func GetById(id int) *UserCourseAnswer {
	record := &UserCourseAnswer{}
	DB.Where("id = ?", id).
		First(&record)
	return record
}

func GetByUserIdByIssueId(user_id, issue_day_id int) *UserCourseAnswer {
	record := &UserCourseAnswer{}
	DB.Where("user_id = ?", user_id).
		Where("issue_day_id = ?", issue_day_id).
		First(&record)
	return record
}

func GetByIssueId(issue_day_id int) []*UserCourseAnswer {
	record := []*UserCourseAnswer{}
	DB.Where("issue_day_id = ?", issue_day_id).
		Find(&record)
	return record
}

func GetByIssueIdNotReview(issue_day_id int) []*UserCourseAnswer {
	record := []*UserCourseAnswer{}
	DB.Where("issue_day_id = ?", issue_day_id).
		Where("is_review = ?", false).
		Find(&record)
	return record
}


func Delete(id int, ctx context.Context) *UserCourseAnswer {
	DB.Where("id = ?", id).Delete(UserCourseAnswer{})
	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackDelete("user_course_answer", id, user_id)

	return GetById(id)
}

func Update(u *UserCourseAnswer, ctx context.Context) *UserCourseAnswer {
	DB.Save(&u)
	u.IsReview = false
	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackEdit("user_course_answer", u.Id, user_id)
	return u
}

func Add(u *UserCourseAnswer, ctx context.Context) *UserCourseAnswer {

	DB.Create(u)

	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackAdd("user_course_answer", u.Id, user_id)

	return u
}

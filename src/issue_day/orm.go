package issue_day

import(
	"context"
)

func GetByDayId(day_id int) []*IssueDay {
	records := []*IssueDay{}
	DB.Where("day_id = ?" , day_id).Find(&records)
	return records
}

func GetByCourseIdDayId(course_id, day int) []*IssueDay {
	records := []*IssueDay{}
	DB.Where("course_id = ?" , course_id).
	Where("day_id = ?" , day).
	Find(&records)
	return records
}

func GetByCourseIdDayIdByUserId(course_id, day int) []*IssueDay {
	records := []*IssueDay{}
	DB.Where("course_id = ?" , course_id).
	Where("day_id = ?" , day).
	Find(&records)
	return records
}

func GetByCourseId(course_id int) []*IssueDay {
	records := []*IssueDay{}
	DB.Where("course_id = ?" , course_id).
	Find(&records)
	return records
}

func GetById(id int) *IssueDay {
	record := &IssueDay{}
	DB.Where("id = ?", id).
		First(&record)
	return record
}

func Delete(id int, ctx context.Context) *IssueDay {
	DB.Where("id = ?", id).Delete(IssueDay{})
	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackDelete("issue_day", id, user_id)

	return GetById(id)
}

func Update(u *IssueDay, ctx context.Context) *IssueDay {
	DB.Save(&u)
	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackEdit("issue_day", u.Id, user_id)
	return u
}

func Add(u *IssueDay, ctx context.Context) *IssueDay {

	DB.Create(u)

	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackAdd("issue_day", u.Id, user_id)

	return u
}

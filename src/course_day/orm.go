package course_day

import(
	"context"
)

func GetByCourseId(course_id int) []*CourseDay {
	records := []*CourseDay{}
	DB.Where("course_id = ?" , course_id).Find(&records)
	return records
}

func GetByCourseIdByDay(course_id, day int) []*CourseDay {
	records := []*CourseDay{}
	DB.Where("course_id = ?" , course_id).
		Where("day = ?" , day).
		Find(&records)
	return records
}

func GetById(id int) *CourseDay {
	record := &CourseDay{}
	DB.Where("id = ?", id).
		First(&record)
	return record
}

func GetByDay(day int) []*CourseDay {
	record := []*CourseDay{}
	DB.Where("day = ?", day).
		Find(&record)
	return record
}

func Delete(id int, ctx context.Context) *CourseDay {
	DB.Where("id = ?", id).Delete(CourseDay{})
	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackDelete("course_day", id, user_id)

	return GetById(id)
}

func Update(u *CourseDay, ctx context.Context) *CourseDay {
	DB.Save(&u)
	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackEdit("course_day", u.Id, user_id)
	return u
}

func Add(u *CourseDay, ctx context.Context) *CourseDay {

	DB.Create(u)

	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackAdd("course_day", u.Id, user_id)

	return u
}

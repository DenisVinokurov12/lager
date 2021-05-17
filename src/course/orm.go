package course

import(
	"context"
)


func All() []*Course {
	records := []*Course{}
	DB.Find(&records)
	return records
}


/*
SELECT `course`.* FROM 
`course` issue_day on issue_day.course_id = course.id user_course on 
user_course.issue_day_id = issue_day.id  GROUP BY user_id
*/
func GetByUserId(user_id int) []*Course {
	record := []*Course{}
	DB.Joins("LEFT JOIN issue_day on issue_day.course_id = course.id").
		Joins("LEFT JOIN user_course on user_course.issue_day_id = issue_day.id").
		Group("user_id, course.id").
		Where("user_course.user_id is not null").
		Where("user_course.user_id = ?", user_id).
		Find(&record)
	return record
}

func Delete(id int, ctx context.Context) *Course {
	DB.Where("id = ?", id).Delete(Course{})

	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackDelete("course", id, user_id)

	return GetById(id)
}


func GetById(id int) *Course {
	record := &Course{}
	DB.Where("id = ?", id).
		First(&record)
	return record
}

func Update(u *Course, ctx context.Context) *Course {
	DB.Save(&u)
	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackEdit("course", u.Id, user_id)
	return u
}

func Add(u *Course, ctx context.Context) *Course {

	DB.Create(u)

	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackAdd("course", u.Id, user_id)

	return u
}

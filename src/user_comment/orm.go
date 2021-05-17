package user_comment

import(
	"context"
)

func GetById(id int) *UserComment {
	record := &UserComment{}
	DB.Where("id = ?", id).
		First(&record)
	return record
}


func GetByIssueId(issue_day_id int) []*UserComment {
	record := []*UserComment{}
	DB.Where("issue_day_id = ?", issue_day_id).
		Find(&record)
	return record
}

func GetNagradaByUserId(user_id int) []string {
	record := []*UserComment{}
	out := []string{}
	DB.Where(`nagrada != ""`).
		Where(`nagrada_user_id = ?` , user_id).
		Find(&record)
	for i := 0; i < len(record); i++ {
		out = append(out , record[i].Nagrada)
	}
	return out
}

func GetByStartCommentId(start_comment_id int) []*UserComment {
	record := []*UserComment{}
	DB.Where("start_comment_id = ?", start_comment_id).
		Find(&record)
	return record
}

func GetByIssueIdByUserId(issue_day_id, user_id int) []*UserComment {
	record := []*UserComment{}
	DB.Where("issue_day_id = ?", issue_day_id).
		Where("user_id = ?", user_id).
		Find(&record)
	return record
}


func GetByIssueIdByUserIdStart0(issue_day_id, user_id int) []*UserComment {
	record := []*UserComment{}
	DB.Where("issue_day_id = ?", issue_day_id).
		Where("user_id = ?", user_id).
		Where("start_comment_id = ?",0 ).
		Find(&record)
	return record
}


func Delete(id int, ctx context.Context) *UserComment {
	DB.Where("id = ?", id).Delete(UserComment{})
	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackDelete("user_comment", id, user_id)

	return GetById(id)
}

func Update(u *UserComment, ctx context.Context) *UserComment {
	DB.Save(&u)
	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackEdit("user_comment", u.Id, user_id)
	return u
}

func Add(u *UserComment, ctx context.Context) *UserComment {

	DB.Create(u)

	user_id := 0
	if v := ctx.Value("init_user"); v != nil {
		user_id = ctx.Value("init_user").(int)
	}
	UserLogCallbackAdd("user_comment", u.Id, user_id)

	return u
}

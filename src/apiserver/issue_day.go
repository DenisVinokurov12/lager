package apiserver


import (
	"github.com/hoisie/web"
	"issue_day"
	"io"
	"context"
	"github.com/google/uuid"
	"strconv"
	"fmt"
	"time"
	"user_course"
	"path/filepath"
	"user"
	"os"
)


func handler_issue_day_by_course_day(ctx *web.Context, course_day_id string, course_id string ) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	course_day_id_int, err := strconv.Atoi(course_day_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	course_id_int, err := strconv.Atoi(course_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	return toJSON(issue_day.GetByCourseIdDayId(course_id_int,course_day_id_int))
}

func handler_user_knowbase_by_course(ctx *web.Context, user_id string, course_id string ) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	user_id_int, err := strconv.Atoi(user_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	course_id_int, err := strconv.Atoi(course_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	all := issue_day.GetKnowbaseByUserByCourseId(user_id_int,course_id_int)

	for i := 0; i < len(all); i++ {
		uc := user_course.GetByIdByUserId(all[i].Id, user_id_int)
		if uc.Id != 0{
			all[i].IsCompleted = uc.IsCompleted
		}
	}


	return toJSON(all)
}

func handler_issue_day_by_course_by_day_user(ctx *web.Context, course_id, course_day_id string ) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}


	course_day_id_int, err := strconv.Atoi(course_day_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	course_id_int, err := strconv.Atoi(course_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	all := issue_day.GetByCourseIdDayId(course_id_int,course_day_id_int)

	for i := 0; i < len(all); i++ {
		uc := user_course.GetByIdByUserId(all[i].Id, init_user.Id)
		if uc.Id != 0{
			all[i].IsCompleted = uc.IsCompleted
		}
	}

	return toJSON(all)
}

func handler_user_issue_completed(ctx *web.Context, issue_day_id, user_id string ) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	issue_day_int, err := strconv.Atoi(issue_day_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	user_id_int, err := strconv.Atoi(user_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	issue_ := issue_day.GetById(issue_day_int)
	user_ := user.GetById(user_id_int)

	if issue_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}
	if user_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	uc := user_course.GetByIdByUserId(issue_.Id, user_.Id)
	if !init_user.IsAccessUpdateUserCourse(uc){
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	if uc.IsCompleted {
		return toJSON(uc)
	}

	uc.IsCompleted = true

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)
	user_course.Update(uc, ctx_)

	if len(issue_day.GetByCourseIdDayIdNotCompleted(issue_.CourseId, issue_.DayId, user_.Id)) == 0 {
		// Ищем слудующий день
		next_day := issue_.DayId + 1
		iss := issue_day.GetByCourseIdDayId(issue_.CourseId, next_day)
		for i := 0; i < len(iss); i++ {
			nuc := &user_course.UserCourse{}
			nuc.UserId = user_.Id
			nuc.IssueDayId = iss[i].Id
			nuc.StartTs = time.Now()
			user_course.Add(nuc, ctx_)
		}
	}

	return toJSON(uc)
}

func handler_issue_day_by_course_by_day_by_user(ctx *web.Context, course_id, course_day_id, user_id string ) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}


	course_day_id_int, err := strconv.Atoi(course_day_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	course_id_int, err := strconv.Atoi(course_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	user_id_int, err := strconv.Atoi(course_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	if !init_user.IsAccessViewUser(user.GetById(user_id_int)) {
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	all := issue_day.GetByCourseIdDayIdByUserId(course_id_int,course_day_id_int)

	for i := 0; i < len(all); i++ {
		uc := user_course.GetByIdByUserId(all[i].Id, user_id_int)
		if uc.Id != 0{
			all[i].IsCompleted = uc.IsCompleted
		}
	}

	return toJSON(all)
}

func handler_issue_day_by_course_user(ctx *web.Context, course_id string ) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	course_id_int, err := strconv.Atoi(course_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	all := issue_day.GetByCourseId(course_id_int)

	for i := 0; i < len(all); i++ {
		uc := user_course.GetByIdByUserId(all[i].Id, init_user.Id)
		if uc.Id != 0{
			all[i].IsCompleted = uc.IsCompleted
		}
	}

	return toJSON(all)
}

func handler_issue_day_by_course_by_user(ctx *web.Context, course_id, user_id string ) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	course_id_int, err := strconv.Atoi(course_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	user_id_int, err := strconv.Atoi(course_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	if !init_user.IsAccessViewUser(user.GetById(user_id_int)) {
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	all := issue_day.GetByCourseId(course_id_int)

	for i := 0; i < len(all); i++ {
		uc := user_course.GetByIdByUserId(all[i].Id, user_id_int)
		if uc.Id != 0{
			all[i].IsCompleted = uc.IsCompleted
		}
	}

	return toJSON(all)
}


func handler_get_issue_day(ctx *web.Context, id string ) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	id_int, err := strconv.Atoi(id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	return toJSON(issue_day.GetById(id_int))
}

func handler_set_issue_day(ctx *web.Context, id string ) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	id_int, err := strconv.Atoi(id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}


	issue_day_ := issue_day.GetById(id_int)

	if id_int == 0 {
		if !init_user.IsAccessAddIssueDay(){
			ctx.ResponseWriter.WriteHeader(403)
			return `{"error" : "forbidden"}`
		}
	}else {
		if id_int != 0 {
			if !init_user.IsAccessUpdateIssueDay(issue_day_){
				ctx.ResponseWriter.WriteHeader(403)
				return `{"error" : "forbidden"}`
			}
		}
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	if id_int == 0 {
		issue_day_ = &issue_day.IssueDay{}
	}

	ctx.Request.ParseMultipartForm(0)

	if string(ctx.Request.FormValue("day_id")) != "" {
		day_id, err := strconv.Atoi(string(ctx.Request.FormValue("day_id")))
		if err == nil {
			issue_day_.DayId = day_id
		}
	}
	if string(ctx.Request.FormValue("course_id")) != "" {
		day_id, err := strconv.Atoi(string(ctx.Request.FormValue("course_id")))
		if err == nil {
			issue_day_.CourseId = day_id
		}
	}
	if string(ctx.Request.FormValue("title")) != "" {
		issue_day_.Title = string(ctx.Request.FormValue("title"))
	}
	if string(ctx.Request.FormValue("description")) != "" {
		issue_day_.Description = string(ctx.Request.FormValue("description"))
	}
	if string(ctx.Request.FormValue("video")) != "" {
		issue_day_.Video = string(ctx.Request.FormValue("video"))
	}
	if string(ctx.Request.FormValue("full_description")) != "" {
		issue_day_.FullDescription = string(ctx.Request.FormValue("full_description"))
	}
	if string(ctx.Request.FormValue("question")) != "" {
		issue_day_.Question = string(ctx.Request.FormValue("question"))
	}

	src, hdr, err := ctx.Request.FormFile("preview_image")
	if err == nil {
		defer src.Close()
		dir_save_avatar := filepath.Join(issue_day.DIR_PREVIEW, hdr.Filename)

		extension := filepath.Ext(dir_save_avatar)
			id := uuid.New()
			dir_save_avatar = filepath.Join(issue_day.DIR_PREVIEW, 
				id.String() + extension)

		dst, err := os.Create(dir_save_avatar)
		if err == nil {
			defer dst.Close()
			io.Copy(dst, src)
			issue_day_.Image = dir_save_avatar
		}
	}
	src, hdr, err = ctx.Request.FormFile("file")
	if err == nil {
		defer src.Close()
		dir_save_avatar := filepath.Join(issue_day.DIR_FILE, hdr.Filename)

		extension := filepath.Ext(dir_save_avatar)
			id := uuid.New()
			dir_save_avatar = filepath.Join(issue_day.DIR_FILE, 
				id.String() + extension)
				
		dst, err := os.Create(dir_save_avatar)
		if err == nil {
			defer dst.Close()
			io.Copy(dst, src)
			issue_day_.File = dir_save_avatar
		}
	}

	if id_int == 0 {
		issue_day.Add(issue_day_, ctx_)
	}else{
		issue_day.Update(issue_day_, ctx_)
	}

	return toJSON(issue_day_)
}

func handler_delete_issue_day(ctx *web.Context, id string ) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}


	id_int, err := strconv.Atoi(id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	c := issue_day.GetById(id_int)

	if !init_user.IsAccessDeleteIssueDay(c){
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	if c.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	issue_day_ := issue_day.Delete(id_int, ctx_)

	if issue_day_.Id == 0 {
		return `{"status" : "deleted"}`
	}

	return `{"status" : "not deleted"}`
}



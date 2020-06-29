package apiserver


import (
	"github.com/hoisie/web"
	"issue_day"
	"context"
	"strconv"
	"fmt"
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

	return toJSON(issue_day.GetByDayId(course_day_id_int))
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
		if init_user.IsAccessAddIssueDay(){
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

	if _, ok := ctx.Params["day_id"]; ok {
		day_id, err := strconv.Atoi(ctx.Params["day_id"])
		if err == nil {
			issue_day_.DayId = day_id
		}
	}

	if _, ok := ctx.Params["title"]; ok {
		issue_day_.Title = ctx.Params["title"]
	}
	if _, ok := ctx.Params["description"]; ok {
		issue_day_.Description = ctx.Params["description"]
	}
	if _, ok := ctx.Params["video"]; ok {
		issue_day_.Video = ctx.Params["video"]
	}
	if _, ok := ctx.Params["full_description"]; ok {
		issue_day_.FullDescription = ctx.Params["full_description"]
	}

	src, hdr, err := ctx.Request.FormFile("preview_image")
	if err == nil {
		defer src.Close()
		dir_save_avatar := filepath.Join(issue_day.DIR_PREVIEW, hdr.Filename)
		dst, err := os.Create(dir_save_avatar)
		if err == nil {
			defer dst.Close()
			issue_day_.Image = dir_save_avatar
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
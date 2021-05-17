package apiserver


import (
	// "encoding/json"
	"github.com/hoisie/web"
	// "io/ioutil"
	"path/filepath"
	"user_course_answer"
	"user_course"
	"github.com/google/uuid"
	"issue_day"
	"io"
	"os"
	// "course"
	"context"
	"strconv"
	"fmt"
)


func handler_issue_answer(ctx *web.Context, id string ) string {

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
	if issue_day_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	// Ищем есть ли у него такое
	user_course_ := user_course.GetById(init_user.Id)
	if user_course_.Id == 0{
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	// Ище ответ который был дан ранее
	user_answer := user_course_answer.GetByUserIdByIssueId(init_user.Id, issue_day_.Id)

	src, hdr, err := ctx.Request.FormFile("answer_file")
	if err == nil {
		defer src.Close()
		dir_save_avatar := filepath.Join(user_course_answer.DIR_ANSWER, hdr.Filename)

		extension := filepath.Ext(dir_save_avatar)
			id := uuid.New()
			dir_save_avatar = filepath.Join(user_course_answer.DIR_ANSWER, 
				id.String() + extension)
				
		dst, err := os.Create(dir_save_avatar)
		if err == nil {
			defer dst.Close()
			io.Copy(dst, src)
			user_answer.File = dir_save_avatar
		}
	}

	ctx.Request.ParseMultipartForm(0)

	if string(ctx.Request.FormValue("answer")) != "" {
		user_answer.Answer = string(ctx.Request.FormValue("answer"))
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	user_course_answer.Add(user_answer, ctx_)
	return toJSON(user_answer)
}

// ВСе ответы на эту задачу
func handler_get_issue_answer(ctx *web.Context, issue_id string ) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, _ := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	issue_id_int, err := strconv.Atoi(issue_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	uca := user_course_answer.GetByIssueId(issue_id_int)
	return toJSON(uca)
}

func handler_get_issue_answer_no_review(ctx *web.Context, issue_id string ) string {
	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, _ := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	issue_id_int, err := strconv.Atoi(issue_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	uca := user_course_answer.GetByIssueIdNotReview(issue_id_int)
	return toJSON(uca)
}

func handler_get_issue_answer_by_user(ctx *web.Context, issue_id, user_id string ) string {
	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, _ := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	issue_id_int, err := strconv.Atoi(issue_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}
	
	user_id_int, err := strconv.Atoi(user_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	uca := user_course_answer.GetByUserIdByIssueId(issue_id_int, user_id_int)
	return toJSON(uca)
}


func handler_issue_answer_put(ctx *web.Context, issue_id, id_comment string ) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	issue_id_int, err := strconv.Atoi(issue_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	id_comment_int, err := strconv.Atoi(id_comment)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	issue_day_ := issue_day.GetById(issue_id_int)
	if issue_day_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	// Ищем есть ли у него такое
	user_course_ := user_course.GetById(init_user.Id)
	if user_course_.Id == 0{
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	// Ище ответ который был дан ранее
	user_answer := user_course_answer.GetById(id_comment_int)

	if user_answer.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	src, hdr, err := ctx.Request.FormFile("answer_file")
	if err == nil {
		defer src.Close()
		dir_save_avatar := filepath.Join(user_course_answer.DIR_ANSWER, hdr.Filename)

		extension := filepath.Ext(dir_save_avatar)
			id := uuid.New()
			dir_save_avatar = filepath.Join(user_course_answer.DIR_ANSWER, 
				id.String() + extension)
				
		dst, err := os.Create(dir_save_avatar)
		if err == nil {
			defer dst.Close()
			io.Copy(dst, src)
			user_answer.File = dir_save_avatar
		}
	}


	ctx.Request.ParseMultipartForm(0)

	if string(ctx.Request.FormValue("answer")) != "" {
		user_answer.Answer = string(ctx.Request.FormValue("answer"))
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	user_course_answer.Update(user_answer, ctx_)
	return toJSON(user_answer)
}
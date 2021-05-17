package apiserver

import (
	"github.com/hoisie/web"
	"path/filepath"
	"user_comment"
	"github.com/google/uuid"
	"issue_day"
	"os"
	"user"
	"context"
	"strconv"
	"user_course"
	"io"
	"fmt"
)

type UserComment struct{
	*user_comment.UserComment 	`json:"user_comment"`
	*user.User 	`json:"user"`
}


func handler_issue_comments(ctx *web.Context, issue_id string ) string {
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

	issue_day_ := issue_day.GetById(issue_id_int)
	if issue_day_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	uc := user_course.GetByIdByUserId(issue_day_.Id, init_user.Id)
	if !init_user.IsAccessViewUserCourse(uc) {
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	ac := user_comment.GetByIssueId(issue_id_int)
	res := []UserComment{}
	for i := 0; i < len(ac); i++ {
		res = append(res , UserComment{
			ac[i],
			user.GetById(ac[i].UserId),
		})
	}

	return toJSON(res) 
}

func handler_issue_comments_by_user(ctx *web.Context, issue_id, user_id string ) string {
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

	user_id_int, err := strconv.Atoi(user_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	issue_day_ := issue_day.GetById(issue_id_int)
	if issue_day_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	user_ := user.GetById(user_id_int)
	if user_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	if !init_user.IsAccessViewUserCommentByUserId(user_) {
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	return toJSON(user_comment.GetByIssueIdByUserId(issue_id_int, user_id_int)) 
}

func handler_issue_comment_add(ctx *web.Context, issue_id string ) string {
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

	issue_day_ := issue_day.GetById(issue_id_int)
	if issue_day_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	uc := &user_comment.UserComment{}
	uc.IssueDayId = issue_day_.Id
	uc.UserId = init_user.Id

	ctx.Request.ParseMultipartForm(0)

	uc.StartCommentId = 0

	if string(ctx.Request.FormValue("comment_id")) != "" {

		r_user_comment, err := strconv.Atoi(ctx.Request.FormValue("comment_id"))
		if err != nil{
			ctx.ResponseWriter.WriteHeader(500)
			return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
		}

		r := user_comment.GetById(r_user_comment)

		if r.StartCommentId == 0 {
			uc.StartCommentId= r.Id
		}else{
			uc.StartCommentId = r.StartCommentId
		}

		uc.UserCommentId = r.Id
	}

	if string(ctx.Request.FormValue("comment")) != "" {
		uc.Comment = string(ctx.Request.FormValue("comment"))
	}

	src, hdr, err := ctx.Request.FormFile("comment_file")
	if err == nil {
		defer src.Close()
		dir_save_avatar := filepath.Join(user_comment.DIR_COMMENT, hdr.Filename)

		extension := filepath.Ext(dir_save_avatar)
			id := uuid.New()
			dir_save_avatar = filepath.Join(user_comment.DIR_COMMENT, 
				id.String() + extension)
				
		dst, err := os.Create(dir_save_avatar)
		if err == nil {
			defer dst.Close()
			io.Copy(dst, src)
			uc.File = dir_save_avatar
		}
	}

	src, hdr, err = ctx.Request.FormFile("nagrada")
	if err == nil {
		defer src.Close()
		if init_user.IsAccessNagarda() {
			dir_save_avatar := filepath.Join(user_comment.DIR_NAGRADA, hdr.Filename)

			extension := filepath.Ext(dir_save_avatar)
			id := uuid.New()
			dir_save_avatar = filepath.Join(user_comment.DIR_NAGRADA, 
				id.String() + extension)

			dst, err := os.Create(dir_save_avatar)
			if err == nil {
				defer dst.Close()
				io.Copy(dst, src)
				uc.Nagrada = dir_save_avatar
			}
		}
	}

	if string(ctx.Request.FormValue("nagrada_user_id")) != "" {
		nagrada_user_id_int, err := strconv.Atoi(ctx.Request.FormValue("nagrada_user_id"))
		if err == nil && init_user.IsAccessNagarda(){
			uc.NagradaUserId = nagrada_user_id_int
		}
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	user_comment.Add(uc, ctx_)
	return toJSON(uc)
}


func handler_issue_comment_put(ctx *web.Context, comment_id string  ) string {
	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	comment_id_int, err := strconv.Atoi(comment_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	uc := user_comment.GetById(comment_id_int)
	if uc.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	ctx.Request.ParseMultipartForm(0)

	if string(ctx.Request.FormValue("comment_id")) != "" {

		r_user_comment, err := strconv.Atoi(string(ctx.Request.FormValue("comment_id")))
		if err != nil{
			ctx.ResponseWriter.WriteHeader(500)
			return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
		}

		r := user_comment.GetById(r_user_comment)
		uc.UserCommentId = r.Id
	}

	if string(ctx.Request.FormValue("comment")) != "" {
		uc.Comment = string(ctx.Request.FormValue("comment"))
	}

	src, hdr, err := ctx.Request.FormFile("comment_file")
	if err == nil {
		defer src.Close()
		dir_save_avatar := filepath.Join(user_comment.DIR_COMMENT, hdr.Filename)

		extension := filepath.Ext(dir_save_avatar)
			id := uuid.New()
			dir_save_avatar = filepath.Join(user_comment.DIR_COMMENT, 
				id.String() + extension)
				
		dst, err := os.Create(dir_save_avatar)
		if err == nil {
			defer dst.Close()
			io.Copy(dst, src)
			uc.File = dir_save_avatar
		}
	}

	src, hdr, err = ctx.Request.FormFile("nagrada")
	if err == nil {
		if init_user.IsAccessNagarda() {
			dir_save_avatar := filepath.Join(user_comment.DIR_NAGRADA, hdr.Filename)

			extension := filepath.Ext(dir_save_avatar)
			id := uuid.New()
			dir_save_avatar = filepath.Join(user_comment.DIR_NAGRADA, 
				id.String() + extension)
				
			dst, err := os.Create(dir_save_avatar)
			if err == nil {
				defer dst.Close()
				io.Copy(dst, src)
				uc.Nagrada = dir_save_avatar
			}
		}
	}

	if string(ctx.Request.FormValue("nagrada_user_id")) != "" {
		nagrada_user_id_int, err := strconv.Atoi(ctx.Request.FormValue("nagrada_user_id"))
		if err == nil && init_user.IsAccessNagarda(){
			uc.NagradaUserId = nagrada_user_id_int
		}
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	user_comment.Update(uc, ctx_)
	return toJSON(uc)
}

func handler_issue_comment_delete(ctx *web.Context, id string ) string {
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

	c := user_comment.GetById(id_int)

	if !init_user.IsAccessDeleteUserComment(c){
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	if c.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	user_comment_ := user_comment.Delete(id_int, ctx_)

	if user_comment_.Id == 0 {
		return `{"status" : "deleted"}`
	}

	return `{"status" : "not deleted"}`
}


func handler_issue_day_all_comment_in_my_comments(ctx *web.Context, issue_day_id, user_id string ) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	issue_day_id_int, err := strconv.Atoi(issue_day_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	user_id_int, err := strconv.Atoi(user_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	user_init_comments := user.GetById(user_id_int)
	if user_init_comments.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	issue_day_ := issue_day.GetById(issue_day_id_int)
	if issue_day_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	if !init_user.IsAccessViewUserCommentByUserId(user_init_comments){
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}


	my_comment_my_issue := user_comment.GetByIssueIdByUserIdStart0(issue_day_id_int, user_id_int)

	res := []UserComment{}

	for i := 0; i < len(my_comment_my_issue); i++ {

		res = append(res , UserComment{
			my_comment_my_issue[i],
			user.GetById(my_comment_my_issue[i].UserId),
		})

		c := user_comment.GetByStartCommentId(my_comment_my_issue[i].Id)
		for j := 0; j < len(c); j++ {
			res = append(res , UserComment{
				c[j],
				user.GetById(c[j].UserId),
			})
		}
	}

	return toJSON(res)
}
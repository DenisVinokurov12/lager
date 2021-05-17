package apiserver

import (
	"github.com/hoisie/web"
	"path/filepath"
	"os"
	"io"
	"course"
	"github.com/google/uuid"
	"context"
	"strconv"
	"fmt"
)


func handler_admin_get_courses(ctx *web.Context) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	return toJSON(course.All())
}


func handler_admin_get_course(ctx *web.Context, id string) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)


	id_int, err := strconv.Atoi(id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	c := course.GetById(id_int)

	return toJSON(c)
}


func handler_admin_post_course(ctx *web.Context, id string) string {

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


	course_ := course.GetById(id_int)

	if id_int == 0 {
		if !init_user.IsAccessAddCourse(){
			ctx.ResponseWriter.WriteHeader(403)
			return `{"error" : "forbidden"}`
		}
	}else {
		if id_int != 0 {
			if !init_user.IsAccessUpdateCourse(course_){
				ctx.ResponseWriter.WriteHeader(403)
				return `{"error" : "forbidden"}`
			}
		}
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	if id_int == 0 {
		course_ = &course.Course{}
	}

	ctx.Request.ParseMultipartForm(0)

	if string(ctx.Request.FormValue("title")) != "" {
		course_.Title = string(ctx.Request.FormValue("title"))
	}

	if string(ctx.Request.FormValue("description")) != "" {
		course_.Description = string(ctx.Request.FormValue("description"))
	}

	src, hdr, err := ctx.Request.FormFile("preview")
	if err == nil {
		defer src.Close()
		dir_save_avatar := filepath.Join(course.DIR_PREVIEW, hdr.Filename)
		extension := filepath.Ext(dir_save_avatar)
			id := uuid.New()
			dir_save_avatar = filepath.Join(course.DIR_PREVIEW, 
				id.String() + extension)
				
		dst, err := os.Create(dir_save_avatar)
		if err == nil {
			defer dst.Close()
			io.Copy(dst, src)
			course_.Image = dir_save_avatar
		}
	}

	if id_int == 0 {
		course.Add(course_, ctx_)
	}else{
		course.Update(course_, ctx_)
	}

	return toJSON(course_)
}


func handler_admin_delete_course(ctx *web.Context, id string) string {

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

	c := course.GetById(id_int)

	if !init_user.IsAccessDeleteCourse(c){
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	if c.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	course_ := course.Delete(id_int, ctx_)

	if course_.Id == 0 {
		return `{"status" : "deleted"}`
	}

	return `{"status" : "not deleted"}`
}
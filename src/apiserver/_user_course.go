package apiserver

import (
	"encoding/json"
	"github.com/hoisie/web"
	"io/ioutil"
	"context"
	"strconv"
	"course_day"
	"user_course"
	"fmt"
)


func handler_set_user_course(ctx *web.Context, id string) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	id_int, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Sprintf(`{"error":"%s"}` , err.Error())
	}

	uc := user_course.GetById(id_int)
	if uc.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	if id_int == 0 {
		if init_user.IsAccessAddUserCourse(){
			ctx.ResponseWriter.WriteHeader(403)
			return `{"error" : "forbidden"}`
		}
	}else {
		if id_int != 0 {
			if !init_user.IsAccessUpdateUserCourse(uc){
				ctx.ResponseWriter.WriteHeader(403)
				return `{"error" : "forbidden"}`
			}
		}
	}

	body, err_body := ioutil.ReadAll(ctx.Request.Body)

	if err_body != nil {
		error_body, _ := json.Marshal(Error{Error: ERROR_INVALID_FORMAT})
		return string(error_body)
	}

	req := UserCourse{}

	err = json.Unmarshal(body, &req)

	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)


	uc.CourseId = req.CourseId
	uc.UserId = req.UserId

	if id_int == 0 {
		user_course.Add(uc, ctx_)
	}else{
		user_course.Update(uc, ctx_)
	}

	return toJSON(uc)
}

func handler_complete_user_course(ctx *web.Context, id string) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	id_int, err := strconv.Atoi(id)
	if err != nil {
		return fmt.Sprintf(`{"error":"%s"}` , err.Error())
	}

	uc := user_course.GetById(id_int)
	if uc.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	if init_user.IsAccessUpdateUserCourse(uc){
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	uc.IsCompleted = true
	user_course.Update(uc, ctx_)

	if len(course_day.GetByDay(uc.Day)) >= len(user_course.GetByDayByCourseSuccess(uc.Day, uc.CourseId, init_user.Id)){
		uc.Day += 1
	}
	user_course.Update(uc, ctx_)

	return toJSON(uc)
}
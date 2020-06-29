package apiserver

import (
	"encoding/json"
	"github.com/hoisie/web"
	"io/ioutil"
	"course"
	"course_day"
	"strconv"
	"context"
	"fmt"
)

func handler_admin_get_course_day(ctx *web.Context, id string) string {

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
		return fmt.Sprintf(`{"error" : "%s"}`, err.Error())
	}

	c := course_day.GetById(id_int)

	if c.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	if !init_user.IsAccessCourseDay(c){
		ctx.ResponseWriter.WriteHeader(500)
		return `{"error" : "forbidden"}`
	}

	return toJSON(c)
}

func handler_admin_get_course_days(ctx *web.Context, course_id string) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	course_id_int, err := strconv.Atoi(course_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(403)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}


	c := course.GetById(course_id_int)

	if c.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	if !init_user.IsAccessCourse(c){
		ctx.ResponseWriter.WriteHeader(500)
		return `{"error" : "forbidden"}`
	}

	course_days_res := []*course_day.CourseDay{}
	course_days := course_day.GetByCourseId(course_id_int)
	for i := 0; i < len(course_days); i++ {
		if !init_user.IsAccessCourseDay(course_days[i]){
			continue
		}
		course_days_res = append(course_days_res, course_days[i])
	}

	return toJSON(course_days_res)
}

func handler_admin_post_course_day(ctx *web.Context, course_day_id string) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}


	course_id_day_int, err := strconv.Atoi(course_day_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(403)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	course_day_ := course_day.GetById(course_id_day_int)
	if course_day_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	body, err_body := ioutil.ReadAll(ctx.Request.Body)

	if err_body != nil {
		error_body, _ := json.Marshal(Error{Error: ERROR_INVALID_FORMAT})
		return string(error_body)
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	req := CourseDayPost{}

	err = json.Unmarshal(body, &req)

	course_day_.Day = req.DayId
	course_day_.CourseId = req.CourseId
	course_day_.Title = req.Title
	course_day_.Description = req.Description

	if course_day_.Id == 0 {
		course_day.Add(course_day_, ctx_)
	}else{
		course_day.Update(course_day_, ctx_)
	}

	return toJSON(course_day.GetById(course_id_day_int))
}

func handler_admin_delete_course_day(ctx *web.Context, course_id string) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}


	course_id_int, err := strconv.Atoi(course_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(403)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	return toJSON(course_day.Delete(course_id_int, ctx_))
}
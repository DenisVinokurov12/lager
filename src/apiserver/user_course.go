package apiserver

import (
	"encoding/json"
	"github.com/hoisie/web"
	"io/ioutil"
	"context"
	"strconv"
	"course_day"
	"issue_day"
	"user"
	"course"
	"user_course"
	"fmt"
	"time"
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


	uc.IssueDayId = req.IssueDayId
	uc.UserId = req.UserId

	if id_int == 0 {
		user_course.Add(uc, ctx_)
	}else{
		user_course.Update(uc, ctx_)
	}

	return toJSON(uc)
}

func handler_set_user_to_course(ctx *web.Context, course_id, user_id string) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	course_id_int, err := strconv.Atoi(course_id)
	if err != nil {
		return fmt.Sprintf(`{"error":"%s"}` , err.Error())
	}

	user_id_int, err := strconv.Atoi(user_id)
	if err != nil {
		return fmt.Sprintf(`{"error":"%s"}` , err.Error())
	}

	course_ := course.GetById(course_id_int)
	if course_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	user_ := user.GetById(user_id_int)
	if user_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}
	
	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	issue_days := issue_day.GetByCourseIdDayId(course_.Id , 1)
	for i := 0; i < len(issue_days); i++ {


		fmt.Printf("%+v\n\n\n", user_course.GetByIdByUserId(issue_days[i].Id, user_.Id))

		if user_course.GetByIdByUserId(issue_days[i].Id, user_.Id).Id == 0 {
			uc := &user_course.UserCourse{
				UserId : user_.Id,
				IssueDayId : issue_days[i].Id,
				StartTs : time.Now(),
			}
			user_course.Add(uc, ctx_)
		}
	}
	return toJSON(user_course.GetByUserIdNotCompleted(user_.Id))
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

	// Получим текущую задачу
	ui := issue_day.GetById(uc.IssueDayId)

	// Получим все уроки на дне
	issues := issue_day.GetByCourseIdDayIdCompleted(ui.CourseId, ui.DayId, uc.UserId)


	if len(course_day.GetByCourseIdByDay(ui.CourseId, ui.DayId)) >= len(issues){
		// uc.Day += 1
		// Ищем слудующий день
		next_day := ui.DayId + 1
		iss := issue_day.GetByCourseIdDayId(ui.CourseId, next_day)
		for i := 0; i < len(iss); i++ {
			nuc := &user_course.UserCourse{}
			nuc.UserId = uc.UserId
			nuc.IssueDayId = iss[i].Id
			nuc.StartTs = time.Now()
			user_course.Add(nuc, ctx_)
		}
	}

	return toJSON(uc)
}
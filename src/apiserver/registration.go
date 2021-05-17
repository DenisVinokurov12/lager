package apiserver

import (
	"encoding/json"
	"github.com/hoisie/web"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"user"
	"time"
	"order"
	"user_course"
	"course_day"
	"context"
)

func handler_registration(ctx *web.Context) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	// is_auth, init_user := is_auth(ctx)

	// fmt.Println(ctx.Request.Header.Get("Origin"))

	// if config.RegistrationHost == ctx.Request.Header.Get("Origin"){
	// 	init_user = &user.User{}
	// }else{
	// 	if !is_auth {
	// 		return `{"error" : "unauthorised"}`
	// 	}
	// }

	body, err_body := ioutil.ReadAll(ctx.Request.Body)

	if err_body != nil {
		error_body, _ := json.Marshal(Error{Error: ERROR_INVALID_FORMAT})
		return string(error_body)
	}

	req := RegistrationPost{}

	err := json.Unmarshal(body, &req)
	if err != nil {
		logrus.Warn("error auth: ", body, " : ", err.Error())
		return `{"error" : "invalid format"}`
	}

	if req.FirstName == "" {
		return `{"error" : "empty first_name"}`
	}
	if req.LastName == "" {
		return `{"error" : "empty last_name"}`
	}
	if req.Phone == "" {
		return `{"error" : "empty phone"}`
	}
	if req.Email == "" {
		return `{"error" : "empty email"}`
	}
	if req.Password == "" {
		return `{"error" : "empty password"}`
	}
	if req.Login == "" {
		return `{"error" : "empty login"}`
	}
	if req.OrderId != "" {
		// Проверим сколько этому 
		order_ := order.GetByOrderId(req.OrderId)
		if order_.Id == 0 {
			return `{"error" : "order_id not found"}`
		}
		if len(user.GetByOrderId(req.OrderId)) > order_.Quantity {
			return `{"error" : "limit user by order_id done"}`
		}

	}

	ctx_ := context.WithValue(context.Background(), "init_user", 0)

	user_ := &user.User{
		FirstName : req.FirstName,
		LastName : req.LastName,
		Phone : req.Phone,
		Login : req.Login,
		Email : req.Email,
		Password : user.CompressPass(req.Password),
		AddTs : time.Now(),
		OrderId : req.OrderId,
		Rule : req.Rule,
		Gender : req.Gender,
	}
	user.Add(user_, ctx_)


	if req.CourseId != 0 && user_course.GetById(req.CourseId).Id != 0 {

		// день который выполнил юзер

		issues_days := course_day.GetByCourseIdByDay(req.CourseId, 1)

		for i := 0; i < len(issues_days); i++ {
			uc := &user_course.UserCourse{}
			uc.UserId = user_.Id
			uc.IssueDayId = issues_days[i].Id
			uc.StartTs = time.Now()
			user_course.Add(uc, ctx_)
		}

	}

	return toJSON(user_)
}
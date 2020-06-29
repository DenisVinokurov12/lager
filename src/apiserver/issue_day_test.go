package apiserver

import (
	"bytes"
	"database/sql"
	// "encoding/json"
	"fmt"
	"github.com/DATA-DOG/go-txdb"
	"github.com/hoisie/web"
	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
	"io"
	"user_course"
	"course_day"
	"context"
	"io/ioutil"
	"net/http"
	"time"
	"course"
	"issue_day"
	"net/http/httptest"
	"testing"
	"user"
)

func init() {
	txdb.Register("device_test",
		"mysql",
		`root:111@/lager_test?charset=utf8&parseTime=True&loc=Local`,
	)
}


// app.Get(`/api/course_day/([0-9]*)/issues`, handler_issue_day_by_course_day)
// 	app.Get(`/api/issues/([0-9]*)`, handler_get_issue_day)
// 	app.Get(`/api/issues/day/([0-9]*)/course/([0-9]*)`, handler_issue_day_by_course_day_by_user)
// 	app.Put(`/api/admin/issues/([0-9]*)`, handler_set_issue_day)
// 	app.Post(`/api/admin/issues/([0-9]*)`, handler_set_issue_day)
// 	app.Delete(`/api/admin/issues/([0-9]*)`, handler_delete_issue_day)


// Получение всех заданий в одном дне на курсе
// func TestIssueDayByCourseDay(t *testing.T) {

// 	s, err := sql.Open("device_test", "device_test_1")
// 	db, err := gorm.Open("mysql", s)

// 	if err != nil {
// 		t.Fatal("Connect to database: ", err.Error())
// 		return
// 	}
// 	defer db.Close()
// 	db.LogMode(false)

// 	course.Init(db)
// 	issue_day.Init(db)
// 	user.Init(db)
// 	user_course.Init(db)
// 	course_day.Init(db)

// 	ctx_ := context.WithValue(context.Background(), "init_user", 0)

// 	course_1 := &course.Course{}
// 	course_1.Title = "Мой первый курс"
// 	course.Add(course_1, ctx_)


// 	course_2 := &course.Course{}
// 	course_2.Title = "Мой второй курс"
// 	course.Add(course_2, ctx_)

// 	issue_day_1_1 := &issue_day.IssueDay{}
// 	issue_day_1_1.DayId = 1
// 	issue_day_1_1.Title = "Курс 1. День 1"
// 	issue_day_1_1.CourseId = course_1.Id
// 	issue_day.Add(issue_day_1_1, ctx_)

// 	issue_day_1_2 := &issue_day.IssueDay{}
// 	issue_day_1_2.DayId = 2
// 	issue_day_1_2.Title = "Курс 1. День 2"
// 	issue_day_1_2.CourseId = course_1.Id
// 	issue_day.Add(issue_day_1_2, ctx_)

// 	issue_day_2_1 := &issue_day.IssueDay{}
// 	issue_day_2_1.DayId = 1
// 	issue_day_2_1.Title = "Курс 2. День 1"
// 	issue_day_2_1.CourseId = course_2.Id
// 	issue_day.Add(issue_day_2_1, ctx_)

// 	issue_day_2_2 := &issue_day.IssueDay{}
// 	issue_day_2_2.DayId = 2
// 	issue_day_2_2.CourseId = course_2.Id
// 	issue_day_2_2.Title = "Курс 2. День 2"
// 	issue_day.Add(issue_day_2_2, ctx_)

// 	new_user := &user.User{
// 		Login:          "MyTestUser",
// 		ApiToken:       "MyApiToken",
// 	}
// 	new_user = user.Add(new_user, ctx_)

// 	sessions.Add("MyApiToken", new_user.Id, cache.NoExpiration)

// 	my_url := func(w http.ResponseWriter, r *http.Request) {
// 		r.Header.Set("Token", "MyApiToken")
// 		ctx := &web.Context{
// 			Request:        r,
// 			Params:         make(map[string]string),
// 			ResponseWriter: w,
// 		}
// 		io.WriteString(w,
// 			handler_issue_day_by_course_day(ctx, "1" , "2"),
// 		)
// 	}

// 	req := httptest.NewRequest("GET", "/", bytes.NewBuffer(make([]byte, 0)))

// 	w := httptest.NewRecorder()
// 	my_url(w, req)

// 	body, _ := ioutil.ReadAll(w.Body)

// 	fmt.Println(string(body))

// }

// Получение всех заданий в одном дне на курсе
func TestIssueDayByCourseDay2(t *testing.T) {

	s, err := sql.Open("device_test", "device_test_1")
	db, err := gorm.Open("mysql", s)

	if err != nil {
		t.Fatal("Connect to database: ", err.Error())
		return
	}
	defer db.Close()
	db.LogMode(true)

	course.Init(db)
	issue_day.Init(db)
	user.Init(db)
	user_course.Init(db)
	course_day.Init(db)

	ctx_ := context.WithValue(context.Background(), "init_user", 0)

	course_1 := &course.Course{}
	course_1.Title = "Мой первый курс"
	course.Add(course_1, ctx_)


	course_2 := &course.Course{}
	course_2.Title = "Мой второй курс"
	course.Add(course_2, ctx_)

	issue_day_1_1 := &issue_day.IssueDay{}
	issue_day_1_1.DayId = 1
	issue_day_1_1.Title = "Курс 1. День 1"
	issue_day_1_1.CourseId = course_1.Id
	issue_day.Add(issue_day_1_1, ctx_)

	issue_day_1_2 := &issue_day.IssueDay{}
	issue_day_1_2.DayId = 2
	issue_day_1_2.Title = "Курс 1. День 2"
	issue_day_1_2.CourseId = course_1.Id
	issue_day.Add(issue_day_1_2, ctx_)

	issue_day_2_1 := &issue_day.IssueDay{}
	issue_day_2_1.DayId = 1
	issue_day_2_1.Title = "Курс 2. День 1"
	issue_day_2_1.CourseId = course_2.Id
	issue_day.Add(issue_day_2_1, ctx_)

	issue_day_2_2 := &issue_day.IssueDay{}
	issue_day_2_2.DayId = 2
	issue_day_2_2.CourseId = course_2.Id
	issue_day_2_2.Title = "Курс 2. День 2"
	issue_day.Add(issue_day_2_2, ctx_)

	new_user := &user.User{
		Login:          "MyTestUser",
		ApiToken:       "MyApiToken",
	}
	new_user = user.Add(new_user, ctx_)
	sessions.Add(new_user.ApiToken, new_user.Id, cache.NoExpiration)

	uc := &user_course.UserCourse{
		UserId : new_user.Id,
		IssueDayId : issue_day_1_2.Id,
		IsCompleted : true,
		StartTs : time.Now(),
	}

	user_course.Add(uc, ctx_)

	my_url := func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Token", new_user.ApiToken)
		ctx := &web.Context{
			Request:        r,
			Params:         make(map[string]string),
			ResponseWriter: w,
		}
		io.WriteString(w,
			handler_issue_day_by_course_user(ctx, 
				fmt.Sprintf(`%d` , course_1.Id)),
		)
	}

	req := httptest.NewRequest("GET", "/", bytes.NewBuffer(make([]byte, 0)))

	w := httptest.NewRecorder()
	my_url(w, req)

	body, _ := ioutil.ReadAll(w.Body)

	fmt.Println(string(body))

}

func TestIssueDayByCourseDay3(t *testing.T) {

	s, err := sql.Open("device_test", "device_test_1")
	db, err := gorm.Open("mysql", s)

	if err != nil {
		t.Fatal("Connect to database: ", err.Error())
		return
	}
	defer db.Close()
	db.LogMode(false)

	course.Init(db)
	issue_day.Init(db)
	user.Init(db)
	user_course.Init(db)
	course_day.Init(db)

	ctx_ := context.WithValue(context.Background(), "init_user", 0)

	course_1 := &course.Course{}
	course_1.Title = "Мой первый курс"
	course.Add(course_1, ctx_)


	course_2 := &course.Course{}
	course_2.Title = "Мой второй курс"
	course.Add(course_2, ctx_)

	issue_day_1_1 := &issue_day.IssueDay{}
	issue_day_1_1.DayId = 1
	issue_day_1_1.Title = "Курс 1. День 1"
	issue_day_1_1.CourseId = course_1.Id
	issue_day.Add(issue_day_1_1, ctx_)

	issue_day_1_2 := &issue_day.IssueDay{}
	issue_day_1_2.DayId = 2
	issue_day_1_2.Title = "Курс 1. День 2"
	issue_day_1_2.CourseId = course_1.Id
	issue_day.Add(issue_day_1_2, ctx_)

	issue_day_2_1 := &issue_day.IssueDay{}
	issue_day_2_1.DayId = 1
	issue_day_2_1.Title = "Курс 2. День 1"
	issue_day_2_1.CourseId = course_2.Id
	issue_day.Add(issue_day_2_1, ctx_)

	issue_day_2_2 := &issue_day.IssueDay{}
	issue_day_2_2.DayId = 2
	issue_day_2_2.CourseId = course_2.Id
	issue_day_2_2.Title = "Курс 2. День 2"
	issue_day.Add(issue_day_2_2, ctx_)

	new_user := &user.User{
		Login:          "MyTestUser",
		ApiToken:       "MyApiToken",
		Rule:      2,
	}
	new_user = user.Add(new_user, ctx_)
	sessions.Add(new_user.ApiToken, new_user.Id, cache.NoExpiration)

	uc := &user_course.UserCourse{
		UserId : new_user.Id,
		IssueDayId : issue_day_1_2.Id,
		IsCompleted : true,
		StartTs : time.Now(),
	}

	user_course.Add(uc, ctx_)

	my_url := func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Token", new_user.ApiToken)
		ctx := &web.Context{
			Request:        r,
			Params:         make(map[string]string),
			ResponseWriter: w,
		}
		io.WriteString(w,
			handler_issue_day_by_course_by_day_user(ctx, 
				fmt.Sprintf(`%d` , course_1.Id), "1"),
		)
	}

	req := httptest.NewRequest("GET", "/", bytes.NewBuffer(make([]byte, 0)))

	w := httptest.NewRecorder()
	my_url(w, req)

	body, _ := ioutil.ReadAll(w.Body)

	fmt.Println(string(body))

}

func TestIssueDayByCourseDay4(t *testing.T) {

	s, err := sql.Open("device_test", "device_test_1")
	db, err := gorm.Open("mysql", s)

	if err != nil {
		t.Fatal("Connect to database: ", err.Error())
		return
	}
	defer db.Close()
	db.LogMode(false)

	course.Init(db)
	issue_day.Init(db)
	user.Init(db)
	user_course.Init(db)
	course_day.Init(db)

	ctx_ := context.WithValue(context.Background(), "init_user", 0)

	course_1 := &course.Course{}
	course_1.Title = "Мой первый курс"
	course.Add(course_1, ctx_)


	course_2 := &course.Course{}
	course_2.Title = "Мой второй курс"
	course.Add(course_2, ctx_)

	issue_day_1_1 := &issue_day.IssueDay{}
	issue_day_1_1.DayId = 1
	issue_day_1_1.Title = "Курс 1. День 1"
	issue_day_1_1.CourseId = course_1.Id
	issue_day.Add(issue_day_1_1, ctx_)

	issue_day_1_2 := &issue_day.IssueDay{}
	issue_day_1_2.DayId = 2
	issue_day_1_2.Title = "Курс 1. День 2"
	issue_day_1_2.CourseId = course_1.Id
	issue_day.Add(issue_day_1_2, ctx_)

	issue_day_2_1 := &issue_day.IssueDay{}
	issue_day_2_1.DayId = 1
	issue_day_2_1.Title = "Курс 2. День 1"
	issue_day_2_1.CourseId = course_2.Id
	issue_day.Add(issue_day_2_1, ctx_)

	issue_day_2_2 := &issue_day.IssueDay{}
	issue_day_2_2.DayId = 2
	issue_day_2_2.CourseId = course_2.Id
	issue_day_2_2.Title = "Курс 2. День 2"
	issue_day.Add(issue_day_2_2, ctx_)

	new_user1 := &user.User{
		Login:          "qwqwq",
	}
	new_user1 = user.Add(new_user1, ctx_)

	new_user := &user.User{
		Login:          "MyTestUser",
		ApiToken:       "MyApiToken",
		// Rule:      2,
	}
	new_user = user.Add(new_user, ctx_)
	sessions.Add(new_user.ApiToken, new_user.Id, cache.NoExpiration)

	uc := &user_course.UserCourse{
		UserId : new_user1.Id,
		IssueDayId : issue_day_1_2.Id,
		IsCompleted : true,
		StartTs : time.Now(),
	}

	user_course.Add(uc, ctx_)

	my_url := func(w http.ResponseWriter, r *http.Request) {
		r.Header.Set("Token", new_user.ApiToken)
		ctx := &web.Context{
			Request:        r,
			Params:         make(map[string]string),
			ResponseWriter: w,
		}
		io.WriteString(w,
			handler_issue_day_by_course_by_user(ctx, 
				fmt.Sprintf(`%d` , course_1.Id), 
				fmt.Sprintf(`%d`, new_user1.Id)),
		)
	}

	req := httptest.NewRequest("GET", "/", bytes.NewBuffer(make([]byte, 0)))

	w := httptest.NewRecorder()
	my_url(w, req)

	body, _ := ioutil.ReadAll(w.Body)

	fmt.Println(string(body))

}

package apiserver

import (
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/hoisie/web"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

var install_server *web.Server

var sessions *cache.Cache
var DiffTime string

var config = &ApiConfig{}

func init() {
	sessions = cache.New(cache.NoExpiration, cache.NoExpiration)
}

func New(c *ApiConfig) {

	config = c

	app := web.NewServer()
	defer app.Close()

	logrus.Info("Инициализация роутов...")

	// auth

	app.Get(`/api/auth`, handler_get_info_by_token)
	app.Post(`/api/auth`, handler_login)
	app.Delete(`/api/auth`, handler_logout)

	app.Post(`/api/registration`, handler_registration)

	app.Get(`/api/courses`, handler_admin_get_courses)
	app.Get(`/api/course/([0-9]*)`, handler_admin_get_course)
	app.Post(`/api/admin/course/([0-9]*)`, handler_admin_post_course)
	app.Put(`/api/admin/course/([0-9]*)`, handler_admin_post_course)
	app.Delete(`/api/admin/course/([0-9]*)`, handler_admin_delete_course)

	app.Get(`/api/admin/course/([0-9]*)/days`, handler_admin_get_course_days)
	app.Get(`/api/admin/course_day/([0-9]*)`, handler_admin_get_course_day)
	app.Post(`/api/admin/course_day/([0-9]*)`, handler_admin_post_course_day)
	app.Put(`/api/admin/course_day/([0-9]*)`, handler_admin_post_course_day)
	app.Delete(`/api/admin/course_day/([0-9]*)`, handler_admin_delete_course_day)


	// Получение заданий по пользователю со статусом
	// Получение заданий на определенном дне у курса
	app.Get(`/api/issues/course/([0-9]*)/day/([0-9]*)/my`, handler_issue_day_by_course_by_day_user)
	app.Get(`/api/user/([0-9]*)/course/([0-9]*)/knowbase`, 
	handler_user_knowbase_by_course)
	

	// Получение цепочки сообщений от своего и до всех кто отвечал
	app.Get(`/api/issues/([0-9]*)/user/([0-9]*)/comments`, 
	handler_issue_day_all_comment_in_my_comments)
	// Получение заданий на курсе
	app.Get(`/api/issues/course/([0-9]*)/my`, handler_issue_day_by_course_user)
	
	app.Get(`/api/issues/course/([0-9]*)/day/([0-9]*)/user/([0-9]*)`, handler_issue_day_by_course_by_day_by_user)

	app.Get(`/api/issues/course/([0-9]*)/user/([0-9]*)`, handler_issue_day_by_course_by_user)
	
	// Ответ пользователя на вопрпос
	// Получение комментариев по задаче
	app.Get(`/api/issues/([0-9]*)/answer`, handler_get_issue_answer)
	app.Get(`/api/issues/([0-9]*)/answer/no_review`, handler_get_issue_answer_no_review)

	// Получение комментариев от какого то пользователя
	app.Get(`/api/issues/([0-9]*)/user/([0-9]*)/answer`, handler_get_issue_answer_by_user)
	// answer_file
	// answer
	// 
	app.Post(`/api/issues/([0-9]*)/answer`, handler_issue_answer)
	app.Put(`/api/issues/([0-9]*)/answer/([0-9]*)`, handler_issue_answer_put)






	// Получение всех комментариев по этой задаче
	app.Get(`/api/issues/([0-9]*)/comments`, handler_issue_comments)
	app.Get(`/api/user/([0-9]*)/get_courses`, handler_get_my_courses)

	// Получение всех комментариев по этой задаче по определенному пользователю
	// app.Get(`/api/issues/([0-9]*)/user/([0-9]*)/comments`, handler_issue_comments_by_user)

	// Добавить комментарий
	app.Post(`/api/issues/([0-9]*)/comment`, handler_issue_comment_add)

	// Ответить на комментарий
	app.Put(`/api/comment/([0-9]*)`, handler_issue_comment_put)

	// Удалить комментарий
	app.Delete(`/api/comment/([0-9]*)`, handler_issue_comment_delete)

	app.Get(`/api/issues/day/([0-9]*)/course/([0-9]*)`, handler_issue_day_by_course_day)
	app.Get(`/api/issues/([0-9]*)`, handler_get_issue_day)
	app.Put(`/api/admin/issues/([0-9]*)`, handler_set_issue_day)
	app.Post(`/api/admin/issues/([0-9]*)`, handler_set_issue_day)
	app.Delete(`/api/admin/issues/([0-9]*)`, handler_delete_issue_day)



	app.Post(`/api/user/update`, handler_user_update)
	app.Post(`/api/order/success`, handler_order_success)
	
	app.Get(`/api/ranks`, handler_get_ranks)

	app.Get(`/api/admin/user_course/([0-9]*)/completed`, handler_complete_user_course)
	app.Get(`/api/admin/issue/([0-9]*)/user/([0-9]*)/completed`, handler_user_issue_completed)
	app.Get(`/api/admin/course/([0-9]*)/user/([0-9]*)/registered`, 
	handler_set_user_to_course)

	app.Post(`/api/admin/user_course/([0-9]*)`, handler_set_user_course)
	app.Put(`/api/admin/user_course/([0-9]*)`, handler_set_user_course)

	
	app.Match("DELETE", `(.*)`, handler_option)
	app.Match("OPTIONS", `(.*)`, handler_option)

	logrus.Infof("Запуск на :%d", c.Port)
	app.Run(fmt.Sprintf(`:%d`, c.Port))

}

func handler_option(ctx *web.Context, url string) string {
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE, OPTIONS", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)
	ctx.SetHeader("Access-Control-Allow-Headers", "Content-Type, Token", true)
	return ``
}

func toJSON(out interface{}) string {
	e, err := json.Marshal(out)
	if err != nil {
		logrus.Fatal("toJSON ", err.Error())
		return `{}`
	}
	return string(e)
}

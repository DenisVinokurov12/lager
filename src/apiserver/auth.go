package apiserver

import (
	"encoding/json"
	"fmt"
	"github.com/hoisie/web"
	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
	"github.com/google/uuid"
	"io/ioutil"
	"time"
	"user"
)

func handler_login(ctx *web.Context) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	body, err_body := ioutil.ReadAll(ctx.Request.Body)

	if err_body != nil {
		error_body, _ := json.Marshal(Error{Error: ERROR_INVALID_FORMAT})
		return string(error_body)
	}

	req := LoginPost{}

	err := json.Unmarshal(body, &req)
	if err != nil {
		logrus.Warn("error auth: ", body, " : ", err.Error())
		return `{"error" : "invalid format"}`
	}

	if req.Username == "" {
		return `{"error" : "empty username"}`
	}

	if req.Password == "" {
		return `{"error" : "empty password"}`
	}

	// Ищем пользователя в БД
	// Если находим то генерим новый токен
	// и пишем через сколько
	out_str_g, _ := uuid.NewUUID()
	out_str := out_str_g.String()

	user_ := user.GetByLoginByPass(req.Username, req.Password)

	if user_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	live := cache.NoExpiration
	if config.SessionLive != 0 {
		live = time.Duration(config.SessionLive) * time.Second
	}

	sessions.Add(out_str, user_.Id, live)

	return fmt.Sprintf(`{"token":"%s", "rule" : "%d"}`, out_str, user_.Rule)
}

func handler_get_info_by_token(ctx *web.Context) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, user_ := is_auth(ctx)

	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	out_str_g, _ := uuid.NewUUID()
	out_str := out_str_g.String()

	for ident, _ := range sessions.Items() {
		id_user, _ := sessions.Get(ident)
		if id_user.(int) == user_.Id {
			sessions.Delete(ident)
		}
	}

	sessions.Add(out_str, user_.Id, cache.NoExpiration)

	return toJSON(struct{
		Token 		string `json:"token"`
		FirstName 	string `json:"first_name"`
		LastName 	string `json:"last_name"`
		Login 		string `json:"login"`
		Email 		string `json:"email"`
		Phone 		string `json:"phone"`
		Avatar 		string `json:"avatar"`
		Rule 		int 	`json:"rule"`
	}{
		Token:    out_str,
		FirstName: user_.FirstName,
		LastName: user_.LastName,
		Login: user_.Login,
		Email: user_.Email,
		Phone: user_.Phone,
		Avatar: user_.Avatar,
		Rule: user_.Rule,
	})

}

func handler_logout(ctx *web.Context) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, _ := is_auth(ctx)

	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	sessions.Delete(ctx.Request.Header["Token"][0])

	return `{"status": "nologged"}`

}

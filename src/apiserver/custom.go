package apiserver

import (
	"github.com/hoisie/web"
	"user"
)

var IS_TESTING_AUTH = false

func is_auth(ctx *web.Context) (bool, *user.User) {

	if IS_TESTING_AUTH {
		return true, user.GetById(1)
	}

	if _, ok := ctx.Request.Header["Token"]; !ok {
		ctx.Unauthorized()
		return false, &user.User{}
	}

	if len(ctx.Request.Header["Token"]) == 0 {
		ctx.Unauthorized()
		return false, &user.User{}
	}

	id_user, foo := sessions.Get(ctx.Request.Header["Token"][0])
	if !foo {
		ctx.Unauthorized()
		return false, &user.User{}
	}

	id_user_int := id_user.(int)

	user_ := user.GetById(id_user_int)
	if user_.Id == 0 {
		ctx.Unauthorized()
		return false, &user.User{}
	}

	// us := UserSession{
	// 	Id : user_.Id,
	// 	Rules : map[string]map[int][]string,
	// }

	return true, user_
}

package apiserver

import (
	// "encoding/json"
	"github.com/hoisie/web"
	// "io/ioutil"
	"course"
	"context"
	"github.com/google/uuid"
	"io"
	"strconv"
	"fmt"
	// "io"
	"os"
	"user"
	"path/filepath"
)

// func handler_admin_get_user_course(ctx *web.Context, id string) string {

// 	ctx.ContentType("json")
// 	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
// 	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

// 	is_auth, _ := is_auth(ctx)
// 	if !is_auth {
// 		return `{"error" : "unauthorised"}`
// 	}

// 	return `sdsdsdsd`

// }

func handler_get_my_courses(ctx *web.Context, user_id string) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	user_id_int, err := strconv.Atoi(user_id)
	if err !=nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	user_ := user.GetById(user_id_int)
	if user_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not found"}`
	}

	if !init_user.IsAccessViewCourseUser(user_){
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	return toJSON(course.GetByUserId(user_id_int))

}


func handler_user_update(ctx *web.Context, user_id string) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	is_auth, init_user := is_auth(ctx)
	if !is_auth {
		return `{"error" : "unauthorised"}`
	}

	user_id_int, err := strconv.Atoi(user_id)
	if err != nil{
		ctx.ResponseWriter.WriteHeader(500)
		return fmt.Sprintf(`{"error" : "%s"}` , err.Error())
	}

	user_ := user.GetById(user_id_int)
	if user_.Id == 0 {
		ctx.ResponseWriter.WriteHeader(404)
		return `{"error" : "not dound"}`
	}

	src, hdr, err := ctx.Request.FormFile("avatar")
	if err == nil {
		defer src.Close()
		dir_save_avatar := filepath.Join(user.DIR_AVATAR, hdr.Filename)

		extension := filepath.Ext(dir_save_avatar)
			id := uuid.New()
			dir_save_avatar = filepath.Join(user.DIR_AVATAR, 
				id.String() + extension)
				
		dst, err := os.Create(dir_save_avatar)
		if err == nil {
			defer dst.Close()
			io.Copy(dst, src)
			init_user.Avatar = dir_save_avatar
		}
	}

	if _, ok := ctx.Params["first_name"]; ok {
		user_.FirstName = ctx.Params["first_name"]
	}

	if _, ok := ctx.Params["last_name"]; ok {
		user_.LastName = ctx.Params["last_name"]
	}

	if _, ok := ctx.Params["login"]; ok {
		user_.Login = ctx.Params["login"]
	}
	if _, ok := ctx.Params["rank_id"]; ok {
		if init_user.IsAccessUpdateRankUser(user_){
			rank_id_int, err := strconv.Atoi(ctx.Params["rank_id"])
			if err == nil {
				user_.RankId = rank_id_int
			}
		}
	}

	if _, ok := ctx.Params["password"]; ok {
		user_.Password = user.CompressPass(ctx.Params["password"])
	}

	if _, ok := ctx.Params["phone"]; ok {
		user_.Phone = user.CompressPass(ctx.Params["phone"])
	}

	if _, ok := ctx.Params["email"]; ok {
		user_.Email = user.CompressPass(ctx.Params["email"])
	}
	if _, ok := ctx.Params["gender"]; ok {
		user_.Gender = user.CompressPass(ctx.Params["gender"])
	}

	ctx_ := context.WithValue(context.Background(), "init_user", init_user.Id)

	if !init_user.IsAccessUpdateUser(user_){
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	user.Update(user_, ctx_)

	return toJSON(user.GetById(user_.Id))
}

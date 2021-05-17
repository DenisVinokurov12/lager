package apiserver
import (
	"github.com/hoisie/web"
	"rank"
)


func handler_get_ranks(ctx *web.Context) string{
	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	return toJSON(rank.All())
}
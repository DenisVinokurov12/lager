package apiserver


import (
	"encoding/json"
	"github.com/hoisie/web"
	"io/ioutil"
	"order"
	"context"
)


func handler_order_success(ctx *web.Context) string {

	ctx.ContentType("json")
	ctx.SetHeader("Access-Control-Allow-Methods", "GET, HEAD, PUT, PATCH, POST, DELETE", true)
	ctx.SetHeader("Access-Control-Allow-Origin", "*", true)

	if config.RegistrationHost != ctx.Request.Header.Get("Origin"){
		ctx.ResponseWriter.WriteHeader(403)
		return `{"error" : "forbidden"}`
	}

	body, err_body := ioutil.ReadAll(ctx.Request.Body)

	if err_body != nil {
		error_body, _ := json.Marshal(Error{Error: ERROR_INVALID_FORMAT})
		return string(error_body)
	}

	req := Order{}

	err := json.Unmarshal(body, &req)
	if err != nil{
		error_body, _ := json.Marshal(Error{Error: err.Error()})
		return string(error_body)
	}

	order_ := order.GetByOrderId(req.OrderId)
	if order_.Id != 0 {
		return `{"error" : "order_id already exists"}`
	}

	ctx_ := context.WithValue(context.Background(), "init_user", 0)

	order_.OrderId = req.OrderId
	order_.Quantity = req.Quantity
	order.Add(order_, ctx_)

	return toJSON(order_)

}

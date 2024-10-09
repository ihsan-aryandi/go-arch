package handler

import (
	"encoding/json"
	"net/http"
)

/*
	Handler
*/

type Handler struct{}

/*
	Context
*/

type Context struct {
	w   http.ResponseWriter
	req *http.Request
}

/*
Types
*/
type Map map[string]interface{}
type Func func(ctx *Context)

func Fn(fn Func) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := &Context{
			w:   writer,
			req: request,
		}

		fn(ctx)
	}
}

func (ctx *Context) ReadBody(v interface{}) error {
	return json.NewDecoder(ctx.req.Body).Decode(v)
}

func (ctx *Context) JSON(statusCode int, data interface{}) {
	ctx.w.WriteHeader(statusCode)

	_ = json.NewEncoder(ctx.w).Encode(data)
}

func (ctx *Context) Request() *http.Request {
	return ctx.req
}

func (ctx *Context) Writer() http.ResponseWriter {
	return ctx.w
}

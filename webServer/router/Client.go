package router

import (
	"fmt"
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/Deansquirrel/goServiceSupport/worker"
	"github.com/kataras/iris"
)

type client struct {
	app *iris.Application
	c   *common
}

func NewRouterClient(app *iris.Application) *client {
	return &client{
		app: app,
		c:   &common{},
	}
}

func (r *client) AddRouter() {
	v := r.app.Party("/client")
	{
		v.Post("/id", r.getClientId)
	}
}

func (r *client) getClientId(ctx iris.Context) {
	var d object.ClientIdRequest
	err := ctx.ReadJSON(&d)
	if err != nil {
		r.c.WriteResponse(ctx, &object.ClientIdResponse{
			ErrCode: -1,
			ErrMsg:  fmt.Sprintf("Bad Request: %s", err.Error()),
		})
		return
	}
	w := worker.NewWorker()
	newId, err := w.GetClientId(d.ClientType, d.HostName, d.DbId, d.DbName)
	if err != nil {
		r.c.WriteResponse(ctx, &object.ClientIdResponse{
			ErrCode: -1,
			ErrMsg:  err.Error(),
		})
		return
	} else {
		r.c.WriteResponse(ctx, &object.ClientIdResponse{
			ErrCode: int(object.ErrTypeCodeNoError),
			ErrMsg:  string(object.ErrTypeMsgNoError),
			Id:      newId,
		})
		return
	}
}

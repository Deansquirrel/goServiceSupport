package router

import (
	"fmt"
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/kataras/iris"
)

import log "github.com/Deansquirrel/goToolLog"

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
	var requestData object.ClientIdRequest
	err := ctx.ReadJSON(&requestData)
	if err != nil {
		log.Debug("")
		r.c.WriteResponse(ctx, &object.ClientIdResponse{
			ErrCode: -1,
			ErrMsg:  fmt.Sprintf("Bad Request: %s", err.Error()),
		})
		return
	}
	//TODO GetClientId
	var clientId string
	r.c.WriteResponse(ctx, &object.ClientIdResponse{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
		Id:      clientId,
	})
}

package router

import (
	"fmt"
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/Deansquirrel/goServiceSupport/repository"
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

	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	idList, err := rep.GetClientId(d.ClientType, d.HostName, d.DbId, d.DbName)
	if err != nil {
		r.c.WriteResponse(ctx, &object.ClientIdResponse{
			ErrCode: -1,
			ErrMsg:  err.Error(),
		})
		return
	}
	if len(idList) > 0 {
		r.c.WriteResponse(ctx, &object.ClientIdResponse{
			ErrCode: int(object.ErrTypeCodeNoError),
			ErrMsg:  string(object.ErrTypeMsgNoError),
			Id:      idList[0],
		})
		return
	}
	newId, err := rep.NewClientId(d.ClientType, d.HostName, d.DbId, d.DbName)
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

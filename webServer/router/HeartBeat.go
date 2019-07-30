package router

import (
	"github.com/Deansquirrel/goServiceSupport/global"
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/kataras/iris"
)

import log "github.com/Deansquirrel/goToolLog"

type heartBeat struct {
	app *iris.Application
	c   *common
}

func NewRouterHeartBeat(app *iris.Application) *heartBeat {
	return &heartBeat{
		app: app,
		c:   &common{},
	}
}

func (r *heartBeat) AddRouter() {
	v := r.app.Party("/heartbeat")
	{
		v.Post("/update", r.update)
	}
}

func (r *heartBeat) update(ctx iris.Context) {
	request := r.c.GetRequestBody(ctx)
	log.Debug(request)
	v := object.TypeResponse{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
		Type:    global.Type,
	}
	r.c.WriteResponse(ctx, v)
}

package router

import (
	"fmt"
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/Deansquirrel/goServiceSupport/worker"
	"github.com/kataras/iris"
)

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
	var d object.HeartBeatRequest
	err := ctx.ReadJSON(&d)
	if err != nil {
		r.c.WriteError(ctx, -1, fmt.Sprintf("Bad Request: %s", err.Error()))
		return
	}
	w := worker.NewWorker()
	err = w.UpdateHeartBeat(&d)
	if err != nil {
		r.c.WriteError(ctx, -1, err.Error())
		return
	}
	r.c.WriteSuccess(ctx)
	//TODO 返回内容增加ClientControl相关，用于控制客户端退出
	return
}

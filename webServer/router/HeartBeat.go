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
	resp := object.HeartBeatResponse{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
	}
	ccList, err := w.GetClientControl(d.ClientId)
	if err != nil {
		r.c.WriteSuccess(ctx)
	} else {
		if len(ccList) > 0 {
			resp.IsForbidden = ccList[0].IsForbidden
			resp.ForbiddenReason = ccList[0].ForbiddenReason
		}
	}
	r.c.WriteResponse(ctx, resp)
	return
}

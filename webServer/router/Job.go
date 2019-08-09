package router

import (
	"fmt"
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/Deansquirrel/goServiceSupport/worker"
	"github.com/kataras/iris"
)

type job struct {
	app *iris.Application
	c   *common
}

func NewRouterJob(app *iris.Application) *job {
	return &job{
		app: app,
		c:   &common{},
	}
}

func (r *job) AddRouter() {
	v := r.app.Party("/job")
	{
		v.Post("/start", r.recordStart)
		v.Post("/end", r.recordEnd)
	}
}

func (r *job) recordStart(ctx iris.Context) {
	var d object.JobRecordRequest
	err := ctx.ReadJSON(&d)
	if err != nil {
		r.c.WriteError(ctx, -1, fmt.Sprintf("Bad Request: %s", err.Error()))
		return
	}
	w := worker.NewWorker()
	err = w.AddJobRecordStart(&d)
	if err != nil {
		r.c.WriteError(ctx, -1, err.Error())
		return
	}
	r.c.WriteSuccess(ctx)
	return
}

func (r *job) recordEnd(ctx iris.Context) {
	var d object.JobRecordRequest
	err := ctx.ReadJSON(&d)
	if err != nil {
		r.c.WriteError(ctx, -1, fmt.Sprintf("Bad Request: %s", err.Error()))
		return
	}
	w := worker.NewWorker()
	err = w.UpdateJobRecordEnd(&d)
	if err != nil {
		r.c.WriteError(ctx, -1, err.Error())
		return
	}
	r.c.WriteSuccess(ctx)
	return
}

//TODO 增加错误记录接口

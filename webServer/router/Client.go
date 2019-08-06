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
		v.Post("/info", r.refreshInfo)
		v.Post("/svrv3", r.refreshSvrV3Info)
	}
}

//func (r *client) getClientId(ctx iris.Context) {
//	var d object.ClientIdRequest
//	err := ctx.ReadJSON(&d)
//	if err != nil {
//		r.c.WriteError(ctx, -1, fmt.Sprintf("Bad Request: %s", err.Error()))
//		return
//	}
//	w := worker.NewWorker()
//	newId, err := w.GetClientId(d.ClientType, d.HostName, d.DbId, d.DbName)
//	if err != nil {
//		r.c.WriteError(ctx, -1, err.Error())
//		return
//	}
//	r.c.WriteResponse(ctx, &object.ClientIdResponse{
//		ErrCode: int(object.ErrTypeCodeNoError),
//		ErrMsg:  string(object.ErrTypeMsgNoError),
//		Id:      newId,
//	})
//	go func() {
//		rList, err := w.GetClientType(d.ClientType)
//		if err == nil && rList != nil && len(rList) < 1 {
//			_ = w.AddNewClientType(d.ClientType)
//		}
//	}()
//	return
//}

func (r *client) refreshInfo(ctx iris.Context) {
	var d object.ClientInfoRequest
	err := ctx.ReadJSON(&d)
	if err != nil {
		r.c.WriteError(ctx, -1, fmt.Sprintf("Bad Request: %s", err.Error()))
		return
	}
	w := worker.NewWorker()
	err = w.RefreshClientInfo(&d)
	if err != nil {
		r.c.WriteError(ctx, -1, err.Error())
		return
	}
	r.c.WriteSuccess(ctx)
	go func() {
		//ClientTypeInfo记录维护（检查新类型插入）
		rList, err := w.GetClientType(d.ClientType)
		if err == nil && rList != nil && len(rList) < 1 {
			_ = w.AddNewClientType(d.ClientType)
		}
	}()
	return
}

func (r *client) refreshSvrV3Info(ctx iris.Context) {
	var d object.SvrV3InfoRequest
	err := ctx.ReadJSON(&d)
	if err != nil {
		r.c.WriteError(ctx, -1, fmt.Sprintf("Bad Request: %s", err.Error()))
		return
	}
	w := worker.NewWorker()
	err = w.RefreshSvrV3Info(&d)
	if err != nil {
		r.c.WriteError(ctx, -1, err.Error())
		return
	}
	r.c.WriteSuccess(ctx)
	return
}

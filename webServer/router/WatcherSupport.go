package router

import (
	"fmt"
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/Deansquirrel/goServiceSupport/worker"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
	"strings"
)

type watcherSupport struct {
	app *iris.Application
	c   *common
}

func NewRouterWatcherSupport(app *iris.Application) *watcherSupport {
	return &watcherSupport{
		app: app,
		c:   &common{},
	}
}

func (r *watcherSupport) AddRouter() {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})
	v := r.app.Party("/watchersupport", crs).AllowMethods(iris.MethodOptions)
	{
		v.Post("/welcome", r.getWelcomeData)
		v.Post("/heartbeatMonitorData", r.getHeartbeatMonitorData)
		v.Post("/delHeartbeat", r.delHeartbeat)
	}
}

func (r *watcherSupport) getWelcomeData(ctx iris.Context) {
	var d object.WelcomeDataRequest
	err := ctx.ReadJSON(&d)
	if err != nil {
		r.c.WriteError(ctx, -1, fmt.Sprintf("Bad Request: %s", err.Error()))
		return
	}

	var typeList []string
	if strings.Trim(d.HeartbeatClientType, " ") != "" {
		typeList = strings.Split(d.HeartbeatClientType, "|")
	}
	w := worker.NewWatcherSupportWorker()
	list, err := w.GetHeartbeatErrCount(typeList)
	if err != nil {
		r.c.WriteError(ctx, -1, err.Error())
		return
	}

	responseData := object.WelcomeDataResponse{
		ErrCode:       int(object.ErrTypeCodeNoError),
		ErrMsg:        string(object.ErrTypeMsgNoError),
		HeartbeatData: list,
	}
	r.c.WriteResponse(ctx, responseData)
	return
}

func (r *watcherSupport) getHeartbeatMonitorData(ctx iris.Context) {
	var d object.HeartbeatMonitorDataRequest
	err := ctx.ReadJSON(&d)
	if err != nil {
		r.c.WriteError(ctx, -1, fmt.Sprintf("Bad Request: %s", err.Error()))
		return
	}
	w := worker.NewWatcherSupportWorker()
	list, err := w.GetHeartbeatMonitorData(d.Type)
	if err != nil {
		r.c.WriteError(ctx, -1, err.Error())
		return
	}
	responseData := object.HeartbeatMonitorDataResponse{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
		Data:    list,
	}
	r.c.WriteResponse(ctx, responseData)
	return
}

func (r *watcherSupport) delHeartbeat(ctx iris.Context) {
	var d object.DelHeartbeatRequest
	err := ctx.ReadJSON(&d)
	if err != nil {
		r.c.WriteError(ctx, -1, fmt.Sprintf("Bad Request: %s", err.Error()))
		return
	}
	w := worker.NewWatcherSupportWorker()
	err = w.DelHeartbeat(d.ClientId)
	if err != nil {
		r.c.WriteError(ctx, -1, err.Error())
		return
	}
	r.c.WriteResponse(ctx, object.Response{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
	})
	return
}

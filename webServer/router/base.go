package router

import (
	"github.com/Deansquirrel/goServiceSupport/global"
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/iris-contrib/middleware/cors"
	"github.com/kataras/iris"
)

type base struct {
	app *iris.Application
	c   common
}

func NewRouterBase(app *iris.Application) *base {
	return &base{
		app: app,
		c:   common{},
	}
}

func (base *base) AddRouter() {
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, //允许通过的主机名称
		AllowCredentials: true,
	})
	v := base.app.Party("/", crs).AllowMethods(iris.MethodOptions)
	//v := base.app.Party("/")
	{
		v.Get("/version", base.version)
		v.Get("/type", base.getType)
	}
}

//获取Type
func (base *base) version(ctx iris.Context) {
	v := object.VersionResponse{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
		Version: global.Version,
	}
	base.c.WriteResponse(ctx, v)
}

//获取版本
func (base *base) getType(ctx iris.Context) {
	v := object.TypeResponse{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
		Type:    global.Type,
	}
	base.c.WriteResponse(ctx, v)
}

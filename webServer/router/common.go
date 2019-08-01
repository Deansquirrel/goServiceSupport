package router

import (
	"encoding/json"
	"fmt"
	"github.com/Deansquirrel/goServiceSupport/object"
	log "github.com/Deansquirrel/goToolLog"
	"github.com/kataras/iris"
	"io/ioutil"
)

const (
	TranErrStr = "{\"errcode\":-1,\"errmsg\":\"构造返回结果时发生错误, %s\"}"
)

type common struct {
}

func (c *common) GetRequestBody(ctx iris.Context) string {
	body := ctx.Request().Body
	defer func() {
		_ = body.Close()
	}()
	b, err := ioutil.ReadAll(body)
	if err != nil {
		log.Error("获取Http请求文本时发生错误：" + err.Error())
		return ""
	}
	return string(b)
}

func (c *common) WriteSuccess(ctx iris.Context) {
	c.WriteResponse(ctx, &object.Response{
		ErrCode: int(object.ErrTypeCodeNoError),
		ErrMsg:  string(object.ErrTypeMsgNoError),
	})
	return
}

func (c *common) WriteError(ctx iris.Context, errCode int, errMsg string) {
	c.WriteResponse(ctx, &object.Response{
		ErrCode: -1,
		ErrMsg:  errMsg,
	})
	return
}

//向ctx中添加返回内容
func (c *common) WriteResponse(ctx iris.Context, v interface{}) {
	str, err := json.Marshal(v)
	if err != nil {
		body := fmt.Sprintf(TranErrStr, "err:"+err.Error())
		_, err = ctx.WriteString(body)
		if err != nil {
			log.Error(fmt.Sprintf("write body err,body: %s,err: %s", string(body), err.Error()))
		}
		return
	}
	_, err = ctx.WriteString(string(str))
	if err != nil {
		log.Error(fmt.Sprintf("write body err,body: %s,err: %s", string(str), err.Error()))
	}
	return
}

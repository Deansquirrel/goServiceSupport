package global

import (
	"context"
	"github.com/Deansquirrel/goServiceSupport/object"
)

const (
	//PreVersion = "1.0.6 Build20190815"
	//TestVersion = "0.0.0 Build20190101"
	Version                   = "1.0.7 Build20190921"
	Type                      = "ServiceSupport"
	SecretKey                 = "ServiceSupport"
	ClearJobRecordCron        = "0 0 * * * ?"
	ClearInvalidHeartBeatCron = "0 0 * * * ?"
)

var Ctx context.Context
var Cancel func()

//程序启动参数
var Args *object.ProgramArgs

//系统参数
var SysConfig *object.SystemConfig

const (
	ListUnknownTitle = "未知"
)

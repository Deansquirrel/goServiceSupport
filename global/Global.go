package global

import (
	"context"
	"github.com/Deansquirrel/goServiceSupport/object"
)

const (
	//PreVersion = "1.0.0 Build20190806"
	//TestVersion = "0.0.0 Build20190101"
	Version                   = "0.0.0 Build20190101"
	Type                      = "ServiceSupport"
	SecretKey                 = "ServiceSupport"
	ClearJobRecordCron        = "5/10 * * * * ?"
	ClearInvalidHeartBeatCron = "0/10 * * * * ?"
)

var Ctx context.Context
var Cancel func()

//程序启动参数
var Args *object.ProgramArgs

//系统参数
var SysConfig *object.SystemConfig

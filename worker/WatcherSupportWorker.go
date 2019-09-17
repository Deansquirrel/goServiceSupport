package worker

import (
	"github.com/Deansquirrel/goServiceSupport/global"
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/Deansquirrel/goServiceSupport/repository"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolMSSql"
	"time"
)

type watcherSupportWorker struct {
	localDbConfig *goToolMSSql.MSSqlConfig
}

func NewWatcherSupportWorker() *watcherSupportWorker {
	return &watcherSupportWorker{
		localDbConfig: repository.NewCommon().GetLocalDbConfig(),
	}
}

func (w *watcherSupportWorker) GetHeartbeatErrCount(typeList []string) ([]object.HeartbeatErrCount, error) {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())

	outTime := time.Now().Add(goToolCommon.GetDurationBySecond(global.SysConfig.SSConfig.HeartBeatForbidden))
	rList, err := rep.GetHeartbeatErrCount(outTime)
	if err != nil {
		return nil, err
	}
	list := make([]object.HeartbeatErrCount, 0)

	//TODO typeList增加位置，然后判断
	for _, d := range rList {

		if d.Type == "未知" || d.Type == "Z5MdDataTrans" || d.Type == "Z9MdDataTransV2" {
			list = append(list, *d)
		}
	}
	return list, nil
}

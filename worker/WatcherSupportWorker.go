package worker

import (
	"github.com/Deansquirrel/goServiceSupport/global"
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/Deansquirrel/goServiceSupport/repository"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolMSSql"
	"sort"
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
	rMap := make(map[string]int)
	if typeList == nil {
		//全部类型
		for _, d := range rList {
			rMap[d.Type] = 0
		}
	} else {
		//指定类型
		for _, cType := range typeList {
			rMap[cType] = 0
		}
	}
	for _, rData := range rList {
		if rData.Type == global.ListUnknownTitle {
			rMap[global.ListUnknownTitle] = rData.Count
			continue
		}
		_, ok := rMap[rData.Type]
		if ok {
			rMap[rData.Type] = rData.Count
		}
	}

	keys := make([]string, 0)
	for k := range rMap {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	list := make([]object.HeartbeatErrCount, 0)
	for _, k := range keys {
		list = append(list, object.HeartbeatErrCount{
			Type:  k,
			Count: rMap[k],
		})
	}
	return list, nil
}

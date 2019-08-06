package worker

import (
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/Deansquirrel/goServiceSupport/repository"
	"github.com/Deansquirrel/goToolMSSql"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"time"
)

type worker struct {
	localDbConfig *goToolMSSql.MSSqlConfig
}

func NewWorker() *worker {
	return &worker{
		localDbConfig: repository.NewCommon().GetLocalDbConfig(),
	}
}

//func (w *worker) GetClientId(clientType string, hostName string, dbId int, dbName string) (string, error) {
//	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
//	idList, err := rep.GetClientId(clientType, hostName, dbId, dbName)
//	if err != nil {
//		return "", err
//	}
//	if len(idList) > 0 {
//		return idList[0], nil
//	}
//	newId, err := rep.NewClientId(clientType, hostName, dbId, dbName)
//	if err != nil {
//		return "", err
//	} else {
//		return newId, nil
//	}
//}

func (w *worker) GetClientType(clientType string) ([]*object.ClientTypeInfo, error) {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	return rep.GetClientType(clientType)
}

func (w *worker) AddNewClientType(clientType string) error {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	return rep.NewClientType(clientType, 0, 0, "")
}

func (w *worker) RefreshClientInfo(d *object.ClientInfoRequest) error {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	return rep.UpdateClientInfo(&object.ClientInfo{
		ClientId:      d.ClientId,
		ClientType:    d.ClientType,
		ClientVersion: d.ClientVersion,
		HostName:      d.HostName,
		DbId:          d.DbId,
		DbName:        d.DbName,
		InternetIP:    d.InternetIP,
		LastUpdate:    time.Now(),
	})
}

func (w *worker) RefreshSvrV3Info(d *object.SvrV3InfoRequest) error {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	return rep.UpdateSvrV3Info(&object.SvrV3Info{
		ClientId:   d.ClientId,
		CoId:       d.CoId,
		CoAb:       d.CoAb,
		CoCode:     d.CoCode,
		CoUserAb:   d.CoUserAb,
		CoUserCode: d.CoUserCode,
		CoFunc:     d.CoFunc,
		SvName:     d.SvName,
		SvVer:      d.SvVer,
		SvDate:     d.SvDate,
		LastUpdate: time.Now(),
	})
}

func (w *worker) UpdateHeartBeat(d *object.HeartBeatRequest) error {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	return rep.UpdateHeartBeat(&object.HeartBeat{
		ClientId:        d.ClientId,
		HeartBeatClient: d.HeartBeatClient,
		HeartBeat:       time.Now(),
	})
}

func (w *worker) AddJobRecordStart(d *object.JobRecordRequest) error {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	return rep.AddJobRecordStart(&object.JobRecord{
		JobId:     d.JobId,
		ClientId:  d.ClientId,
		JobKey:    d.JobKey,
		StartTime: time.Now(),
		EndTime:   goToolMSSqlHelper.GetDefaultOprTime(),
	})
}

func (w *worker) UpdateJobRecordEnd(d *object.JobRecordRequest) error {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	return rep.UpdateJobRecordEnd(&object.JobRecord{
		JobId:     d.JobId,
		ClientId:  d.ClientId,
		JobKey:    d.JobKey,
		StartTime: goToolMSSqlHelper.GetDefaultOprTime(),
		EndTime:   time.Now(),
	})
}

func (w *worker) GetClientControl(id string) ([]*object.ClientControl, error) {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	return rep.GetClientControl(id)
}

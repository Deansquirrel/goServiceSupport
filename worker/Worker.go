package worker

import (
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/Deansquirrel/goServiceSupport/repository"
	"github.com/Deansquirrel/goToolMSSql"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"github.com/kataras/iris/core/errors"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

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

func (w *worker) RefreshSvrZ5ZlVersion(d *object.SvrZ5ZlVersionRequest) error {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	return rep.UpdateSvrZ5ZlVersion(&object.SvrZ5ZlVersion{
		ClientId:      d.ClientId,
		ObjectName:    d.ObjectName,
		ObjectType:    d.ObjectType,
		ObjectVersion: d.ObjectVersion,
		ObjectDate:    d.ObjectDate,
		LastUpdate:    time.Now(),
	})
}

func (w *worker) RefreshSvrZ5ZlCompany(d *object.SvrZ5ZlCompanyRequest) error {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	return rep.UpdateSvrZ5ZlCompany(&object.SvrZ5ZlCompany{
		ClientId:    d.ClientId,
		CoId:        d.CoId,
		CoAb:        d.CoAb,
		CoCode:      d.CoCode,
		CoType:      d.CoType,
		CoUserAb:    d.CoUserAb,
		CoUserCode:  d.CoUserCode,
		CoAccCrDate: d.CoAccCrDate,
		LastUpdate:  time.Now(),
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

func (w *worker) ClearJobRecord() {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	_ = rep.ClearJobRecord()
}

func (w *worker) ClearInvalidHeartBeat() {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	_ = rep.ClearInvalidHeartBeat()
}

func (w *worker) AddJobErrRecord(d *object.JobErrRecordRequest) error {
	if d == nil {
		errMsg := "AddJobErrRecord request is nil"
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	recordList := object.GetJobErrRecordByRequest(d, 1000)
	if recordList == nil {
		errMsg := "GetJobErrRecordByRequest return is nil"
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	if len(recordList) > 0 {
		for _, d := range recordList {
			err := rep.AddJobErrRecord(d)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

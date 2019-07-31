package worker

import (
	"github.com/Deansquirrel/goServiceSupport/repository"
	"github.com/Deansquirrel/goToolMSSql"
)

type worker struct {
	localDbConfig *goToolMSSql.MSSqlConfig
}

func NewWorker() *worker {
	return &worker{
		localDbConfig: repository.NewCommon().GetLocalDbConfig(),
	}
}

func (w *worker) GetClientId(clientType string, hostName string, dbId int, dbName string) (string, error) {
	rep := repository.NewRepLocal(repository.NewCommon().GetLocalDbConfig())
	idList, err := rep.GetClientId(clientType, hostName, dbId, dbName)
	if err != nil {
		return "", err
	}
	if len(idList) > 0 {
		return idList[0], nil
	}
	newId, err := rep.NewClientId(clientType, hostName, dbId, dbName)
	if err != nil {
		return "", err
	} else {
		return newId, nil
	}
}

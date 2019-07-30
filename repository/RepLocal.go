package repository

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goToolCommon"
	"github.com/Deansquirrel/goToolMSSql"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"strings"
)

import log "github.com/Deansquirrel/goToolLog"

const (
	sqlNewClientId = "" +
		"INSERT INTO [clientinfo]([clientid],[clienttype],[hostname],[dbid],[dbname],[lastupdate]) " +
		"VALUES (?,?,?,?,?,GETDATE())"
	sqlGetClientId = "" +
		"SELECT [clientid] " +
		"FROM [clientinfo] " +
		"WHERE 1=1 " +
		"	AND [clienttype] = ? " +
		"	AND [hostname] = ? " +
		"	AND [dbid] = ? " +
		"	AND [dbname] = ?"
)

type repLocal struct {
	dbConfig *goToolMSSql.MSSqlConfig
}

func NewRepLocal(config *goToolMSSql.MSSqlConfig) *repLocal {
	return &repLocal{
		dbConfig: config,
	}
}

func (r *repLocal) NewClientId(clientType, hostName string, dbId int, dbName string) (string, error) {
	newId := r.newClientId()
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlNewClientId,
		newId, clientType, hostName, dbId, dbName)
	if err != nil {
		errMsg := fmt.Sprintf("create new client id err: %s", err.Error())
		log.Error(errMsg)
		return "", errors.New(errMsg)
	}
	return newId, nil
}

func (r *repLocal) newClientId() string {
	id := goToolCommon.Guid()
	return strings.Replace(id, "-", "", -1)
}

func (r *repLocal) GetClientId(clientType, hostName string, dbId int, dbName string) ([]string, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL(r.dbConfig, sqlGetClientId,
		clientType, hostName, dbId, dbName)
	if err != nil {
		errMsg := fmt.Sprintf("get client id err: %s", err.Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	defer func() {
		_ = rows.Close()
	}()
	rList := make([]string, 0)
	for rows.Next() {
		var id string
		err := rows.Scan(&id)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetClientId", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, id)
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetClientId", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

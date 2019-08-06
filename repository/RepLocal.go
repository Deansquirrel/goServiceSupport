package repository

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/Deansquirrel/goToolMSSql"
	"github.com/Deansquirrel/goToolMSSqlHelper"
	"time"
)

import log "github.com/Deansquirrel/goToolLog"

const (
	//sqlNewClientId = "" +
	//	"INSERT INTO [clientinfo]([clientid],[clienttype],[hostname],[dbid],[dbname],[lastupdate]) " +
	//	"VALUES (?,?,?,?,?,GETDATE())"
	//sqlGetClientId = "" +
	//	"SELECT [clientid] " +
	//	"FROM [clientinfo] " +
	//	"WHERE 1=1 " +
	//	"	AND [clienttype] = ? " +
	//	"	AND [hostname] = ? " +
	//	"	AND [dbid] = ? " +
	//	"	AND [dbname] = ?"
	sqlGetClientType = "" +
		"SELECT [clienttype],[issvrv3],[hasdb],[lastversion] " +
		"FROM [clienttypeinfo] " +
		"WHERE [clienttype]=?"
	sqlNewClientType = "" +
		"INSERT INTO [clienttypeinfo]([clienttype],[issvrv3],[hasdb],[lastversion]) " +
		"VALUES (?,?,?,?)"
	sqlUpdateSvrV3Info = "" +
		"IF EXISTS (SELECT * FROM [SvrV3Info] WHERE [clientid] = ?) " +
		"	Begin " +
		"		UPDATE [SvrV3Info] " +
		"		SET [clientId]=?,[coid]=?,[coab]=?,[cocode]=?,[couserab]=?, " +
		"			[cousercode]=?,[cofunc]=?,[svname]=?,[svver]=?,[svdate]=?, " +
		"			[lastupdate]=? " +
		"		WHERE [clientid] = ? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		INSERT INTO [SvrV3Info]([clientId],[coid],[coab],[cocode],[couserab], " +
		"			[cousercode],[cofunc],[svname],[svver],[svdate], " +
		"			[lastupdate]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?,?, " +
		"			?) " +
		"	End"
	sqlUpdateHeartBeat = "" +
		"IF EXISTS (SELECT * FROM [HeartBeat] WHERE [clientid] = ?) " +
		"	Begin " +
		"		UPDATE [HeartBeat] " +
		"		SET [clientId]=?,[heartbeatClient]=?,[heartbeat]=? " +
		"		WHERE [clientid] = ? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		INSERT INTO [HeartBeat]([clientId],[heartbeatClient],[heartbeat]) " +
		"		VALUES (?,?,?) " +
		"	End"
	sqlAddJobRecordStartInfo = "" +
		"INSERT INTO [JobRecord]([jobid],[clientId],[jobkey],[starttime],[endtime]) " +
		"VALUES (?,?,?,?,?)"
	sqlUpdateJobRecordEndInfo = "" +
		"IF EXISTS (SELECT * FROM [JobRecord] WHERE [jobid] = ?) " +
		"	Begin " +
		"		UPDATE [JobRecord] " +
		"		SET [endtime] = ? " +
		"		WHERE [jobid] = ? " +
		"	End " +
		"Else " +
		"	Begin " +
		"		INSERT INTO [JobRecord]([jobid],[clientId],[jobkey],[starttime],[endtime]) " +
		"		VALUES (?,?,?,?,?) " +
		"	End"

	sqlUpdateClientInfo = "" +
		"IF EXISTS (SELECT * FROM [clientinfo] WHERE [clientid] = ?) " +
		"	Begin " +
		"		UPDATE [clientinfo] " +
		"		SET [clientid]=?,[clienttype]=?,[clientversion]=?,[hostname]=?,[dbid]=?, " +
		"			[dbname]=?,[internetip]=?,[lastupdate]=? " +
		"		WHERE [clientid] = ? " +
		"	End " +
		"Else " +
		"	Begin " +
		"		INSERT INTO [clientinfo]([clientid],[clienttype],[clientversion],[hostname],[dbid], " +
		"			[dbname],[internetip],[lastupdate]) " +
		"		VALUES (?,?,?,?,?," +
		"			?,?,?) " +
		"	End"

	sqlGetClientControl = "" +
		"SELECT [clientid],[isforbidden],[forbiddenreason],[lastupdate] " +
		"FROM [clientcontrol] " +
		"WHERE [clientid] = ?"
)

type repLocal struct {
	dbConfig *goToolMSSql.MSSqlConfig
}

func NewRepLocal(config *goToolMSSql.MSSqlConfig) *repLocal {
	return &repLocal{
		dbConfig: config,
	}
}

//func (r *repLocal) NewClientId(clientType, hostName string, dbId int, dbName string) (string, error) {
//	newId := r.newClientId()
//	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlNewClientId,
//		newId, clientType, hostName, dbId, dbName)
//	if err != nil {
//		errMsg := fmt.Sprintf("create new client id err: %s", err.Error())
//		log.Error(errMsg)
//		return "", errors.New(errMsg)
//	}
//	return newId, nil
//}

//func (r *repLocal) newClientId() string {
//	id := goToolCommon.Guid()
//	return strings.Replace(id, "-", "", -1)
//}

//func (r *repLocal) GetClientId(clientType, hostName string, dbId int, dbName string) ([]string, error) {
//	rows, err := goToolMSSqlHelper.GetRowsBySQL(r.dbConfig, sqlGetClientId,
//		clientType, hostName, dbId, dbName)
//	if err != nil {
//		errMsg := fmt.Sprintf("get client id err: %s", err.Error())
//		log.Error(errMsg)
//		return nil, errors.New(errMsg)
//	}
//	defer func() {
//		_ = rows.Close()
//	}()
//	rList := make([]string, 0)
//	for rows.Next() {
//		var id string
//		err := rows.Scan(&id)
//		if err != nil {
//			errMsg := fmt.Sprintf("%s read data err: %s", "GetClientId", err.Error())
//			log.Error(errMsg)
//			return nil, errors.New(errMsg)
//		}
//		rList = append(rList, id)
//	}
//	if rows.Err() != nil {
//		errMsg := fmt.Sprintf("%s read data err: %s", "GetClientId", rows.Err().Error())
//		log.Error(errMsg)
//		return nil, errors.New(errMsg)
//	}
//	return rList, nil
//}

func (r *repLocal) GetClientType(id string) ([]*object.ClientTypeInfo, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL(r.dbConfig, sqlGetClientType, id)
	if err != nil {
		errMsg := fmt.Sprintf("get client type %s err: %s", id, err.Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	defer func() {
		_ = rows.Close()
	}()
	rList := make([]*object.ClientTypeInfo, 0)
	var clientType, lastVersion string
	var isSvrV3, hasDb int
	for rows.Next() {
		err := rows.Scan(&clientType, &isSvrV3, &hasDb, &lastVersion)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetClientType", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.ClientTypeInfo{
			ClientType:  clientType,
			IsSvrV3:     isSvrV3,
			HasDb:       hasDb,
			LastVersion: lastVersion,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetClientType", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

func (r *repLocal) NewClientType(clientType string, isSvrV3 int, hasDb int, lastVersion string) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlNewClientType,
		clientType, isSvrV3, hasDb, lastVersion)
	if err != nil {
		errMsg := fmt.Sprintf("add new client type err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (r *repLocal) UpdateClientInfo(d *object.ClientInfo) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateClientInfo,
		d.ClientId,
		d.ClientId, d.ClientType, d.ClientVersion, d.HostName, d.DbId,
		d.DbName, d.InternetIP, d.LastUpdate,
		d.ClientId,
		d.ClientId, d.ClientType, d.ClientVersion, d.HostName, d.DbId,
		d.DbName, d.InternetIP, d.LastUpdate)
	if err != nil {
		errMsg := fmt.Sprintf("UpdateClientInfo err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (r *repLocal) UpdateSvrV3Info(d *object.SvrV3Info) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateSvrV3Info,
		d.ClientId,
		d.ClientId, d.CoId, d.CoAb, d.CoCode, d.CoUserAb,
		d.CoUserCode, d.CoFunc, d.SvName, d.SvVer, d.SvDate,
		d.LastUpdate,
		d.ClientId,
		d.ClientId, d.CoId, d.CoAb, d.CoCode, d.CoUserAb,
		d.CoUserCode, d.CoFunc, d.SvName, d.SvVer, d.SvDate,
		d.LastUpdate)
	if err != nil {
		errMsg := fmt.Sprintf("UpdateSvrV3Info err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (r *repLocal) UpdateHeartBeat(d *object.HeartBeat) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateHeartBeat,
		d.ClientId,
		d.ClientId, d.HeartBeatClient, d.HeartBeat,
		d.ClientId,
		d.ClientId, d.HeartBeatClient, d.HeartBeat)
	if err != nil {
		errMsg := fmt.Sprintf("UpdateHeartBeat err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (r *repLocal) AddJobRecordStart(d *object.JobRecord) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlAddJobRecordStartInfo,
		d.JobId, d.ClientId, d.JobKey, d.StartTime, d.EndTime)
	if err != nil {
		errMsg := fmt.Sprintf("AddJobRecordStart err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (r *repLocal) UpdateJobRecordEnd(d *object.JobRecord) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateJobRecordEndInfo,
		d.JobId,
		d.EndTime, d.JobId,
		d.JobId, d.ClientId, d.JobKey, d.StartTime, d.EndTime)
	if err != nil {
		errMsg := fmt.Sprintf("UpdateJobRecordEnd err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (r *repLocal) GetClientControl(id string) ([]*object.ClientControl, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL(r.dbConfig, sqlGetClientControl, id)
	if err != nil {
		errMsg := fmt.Sprintf("get client control %s err: %s", id, err.Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	defer func() {
		_ = rows.Close()
	}()
	rList := make([]*object.ClientControl, 0)
	var clientId, forbiddenReason string
	var isForbidden int
	var lastUpdate time.Time
	for rows.Next() {
		err := rows.Scan(&clientId, &isForbidden, &forbiddenReason, &lastUpdate)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetClientControl", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.ClientControl{
			ClientId:        clientId,
			IsForbidden:     isForbidden,
			ForbiddenReason: forbiddenReason,
			LastUpdate:      lastUpdate,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("%s read data err: %s", "GetClientControl", rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

//TODO 获取ClientControl，用于心跳返回，控制客户端退出
//TODO 定期删除JobRecord
//TODO 定期删除无效心跳
//TODO ClientControl内容维护（界面）

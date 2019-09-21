package repository

import (
	"errors"
	"fmt"
	"github.com/Deansquirrel/goServiceSupport/global"
	"github.com/Deansquirrel/goServiceSupport/object"
	"github.com/Deansquirrel/goToolCommon"
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
		"IF NOT EXISTS (SELECT * FROM [clienttypeinfo] WHERE [clienttype] = ?) " +
		"	Begin " +
		"		INSERT INTO [clienttypeinfo]([clienttype],[issvrv3],[hasdb],[lastversion]) " +
		"		VALUES (?,?,?,?) " +
		"	End"
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
	sqlUpdateSvrZ5ZlVersion = "" +
		"IF EXISTS (SELECT * FROM [svrz5zlversion] WHERE [clientId]=? AND [objectName]=?) " +
		"	Begin " +
		"		UPDATE [svrz5zlversion] " +
		"		SET [clientId]=?,[objectName]=?,[objectType]=?,[objectVersion]=?,[objectDate]=?," +
		"			[lastupdate]=? " +
		"		WHERE [clientId]=? AND [objectName]=? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		INSERT INTO [svrz5zlversion]([clientId],[objectName],[objectType],[objectVersion],[objectDate]," +
		"				[lastupdate]) " +
		"		VALUES (" +
		"			?,?,?,?,?," +
		"			?) " +
		"	End"
	sqlUpdateSvrZ5ZlCompany = "" +
		"IF EXISTS (SELECT * FROM [svrz5zlcompany] WHERE [clientId]=?) " +
		"	Begin " +
		"		UPDATE [svrz5zlcompany] " +
		"		SET [clientId]=?,[coId]=?,[coAb]=?,[coCode]=?,[coType]=?, " +
		"			[coUserAb]=?,[coUserCode]=?,[coAccCrDate]=?,[lastUpdate]=? " +
		"		WHERE [clientId]=? " +
		"	End " +
		"ELSE " +
		"	Begin " +
		"		INSERT INTO [svrz5zlcompany]([clientId],[coId],[coAb],[coCode],[coType], " +
		"			[coUserAb],[coUserCode],[coAccCrDate],[lastUpdate]) " +
		"		VALUES ( " +
		"			?,?,?,?,?, " +
		"			?,?,?,?) " +
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

	sqlClearJobRecord = "" +
		"DELETE FROM [jobrecord] " +
		"WHERE [starttime] < ? or [endtime] < ?"

	sqlClearInvalidHeartBeat = "" +
		"DELETE FROM [heartbeat] " +
		"WHERE [heartbeat] < ?"

	sqlAddJobErrRecord = "" +
		"INSERT INTO  [joberrrecord]([jobid],[errhh],[errmsg],[occurclienttime],[occurtime]) " +
		"VALUES(?,?,?,?,?)"

	sqlGetHeartbeatErrCount = "" +
		"select ISNULL(b.clienttype,'" + global.ListUnknownTitle + "') as clienttype,count(a.clientid) as num " +
		"from heartbeat a " +
		"left join clientinfo b on a.clientid = b.clientid " +
		"where a.heartbeat <= ? " +
		"group by b.clienttype " +
		"order by b.clienttype"

	sqlGetHeartbeatMonitorData = "" +
		"select b.clientid, " +
		"	ISNULL(a.coid,-1) AS coid, " +
		"	ISNULL(a.coab,'" + global.ListUnknownTitle + "') as coab, " +
		"	ISNULL(a.couserab,'" + global.ListUnknownTitle + "') AS couserab, " +
		"	ISNULL(a.svver,'" + global.ListUnknownTitle + "') as svver, " +
		"	ISNULL(b.heartbeat,'1900-01-01') as heartbeat, " +
		"	ISNULL(c.clientversion,'" + global.ListUnknownTitle + "') as clientversion " +
		"from heartbeat b " +
		"left join ( " +
		"	SELECT top 0 '' AS clientid,0 as coid,'' as coab,'' as couserab,'' as svver " +
		"	union all " +
		"	select clientid,coid,coab,couserab,svver " +
		"	from svrv3info " +
		"	union all " +
		"	select a.clientid,coid,coab,couserab,ISNULL(b.objectversion,'wz') " +
		"	from svrz5zlcompany a " +
		"	left join svrz5zlversion b on a.clientid = b.clientid and b.objectname = '' " +
		"	) a on a.clientid = b.clientid " +
		"left join clientinfo c on a.clientid = c.clientid " +
		"where c.clienttype = ?"

	sqlDelHeartbeat = "" +
		"delete from heartbeat " +
		"where clientid = ?"

//	sqlS = "" +
//		"SELECT top 0 '' AS clientid,0 as coid,'' as coab,'' as couserab,'' as svver
//	union all
//select clientid,coid,coab,couserab,svver
//from svrv3info
//union all
//select a.clientid,coid,coab,couserab,ISNULL(b.objectversion,'')
//from svrz5zlcompany a
//left join svrz5zlversion b on a.clientid = b.clientid and b.objectname = ''
//
//--select * from svrz5zlcompany
//--select * from svrz5zlversion order by clientid
//
//
//SELECT *
//FROM heartbeat"
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
		clientType,
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

func (r *repLocal) UpdateSvrZ5ZlVersion(d *object.SvrZ5ZlVersion) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateSvrZ5ZlVersion,
		d.ClientId, d.ObjectName,
		d.ClientId, d.ObjectName, d.ObjectType, d.ObjectVersion, d.ObjectDate,
		d.LastUpdate,
		d.ClientId, d.ObjectName,
		d.ClientId, d.ObjectName, d.ObjectType, d.ObjectVersion, d.ObjectDate,
		d.LastUpdate)
	if err != nil {
		errMsg := fmt.Sprintf("UpdateSvrZ5ZlVersion err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

func (r *repLocal) UpdateSvrZ5ZlCompany(d *object.SvrZ5ZlCompany) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlUpdateSvrZ5ZlCompany,
		d.ClientId,
		d.ClientId, d.CoId, d.CoAb, d.CoCode, d.CoType,
		d.CoUserAb, d.CoUserCode, d.CoAccCrDate, d.LastUpdate,
		d.ClientId,
		d.ClientId, d.CoId, d.CoAb, d.CoCode, d.CoType,
		d.CoUserAb, d.CoUserCode, d.CoAccCrDate, d.LastUpdate)
	if err != nil {
		errMsg := fmt.Sprintf("UpdateSvrZ5ZlCompany err: %s", err.Error())
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

func (r *repLocal) AddJobErrRecord(d *object.JobErrRecord) error {
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlAddJobErrRecord,
		d.JobId, d.ErrHh, d.ErrMsg, d.OccurClientTime, d.OccurTime)
	if err != nil {
		errMsg := fmt.Sprintf("AddJobErrRecord err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

//定期删除JobRecord
func (r *repLocal) ClearJobRecord() error {
	t := time.Now().Add(-goToolCommon.GetDurationByDay(global.SysConfig.SSConfig.SaveJobRecord))
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlClearJobRecord,
		goToolCommon.GetDateTimeStrWithMillisecond(t),
		goToolCommon.GetDateTimeStrWithMillisecond(t))
	if err != nil {
		errMsg := fmt.Sprintf("ClearJobRecord err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

//TODO 定期清除错误记录

//定期删除无效心跳
func (r *repLocal) ClearInvalidHeartBeat() error {
	t := time.Now().Add(-goToolCommon.GetDurationByDay(global.SysConfig.SSConfig.SaveForbiddenHeartBeat))
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlClearInvalidHeartBeat,
		goToolCommon.GetDateTimeStrWithMillisecond(t))
	if err != nil {
		errMsg := fmt.Sprintf("ClearInvalidHeartBeat err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

//TODO ClientControl内容维护（界面）

func (r *repLocal) GetHeartbeatErrCount(t time.Time) ([]*object.HeartbeatErrCount, error) {
	outTime := goToolCommon.GetDateTimeStrWithMillisecond(t)
	rows, err := goToolMSSqlHelper.GetRowsBySQL(r.dbConfig, sqlGetHeartbeatErrCount, outTime)
	if err != nil {
		errMsg := fmt.Sprintf("GetHeartbeatErrCount %s err: %s", outTime, err.Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	defer func() {
		_ = rows.Close()
	}()
	rList := make([]*object.HeartbeatErrCount, 0)
	var clientType string
	var count int
	for rows.Next() {
		err := rows.Scan(&clientType, &count)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetClientControl", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		rList = append(rList, &object.HeartbeatErrCount{
			Type:  clientType,
			Count: count,
		})
	}
	if rows.Err() != nil {
		errMsg := fmt.Sprintf("GetHeartbeatErrCount %s err: %s", outTime, rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

func (r *repLocal) GetHeartbeatMonitorData(cType string) ([]*object.HeartbeatMonitorData, error) {
	rows, err := goToolMSSqlHelper.GetRowsBySQL(r.dbConfig, sqlGetHeartbeatMonitorData, cType)
	if err != nil {
		errMsg := fmt.Sprintf("GetHeartbeatMonitorData %s err: %s", cType, err.Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	defer func() {
		_ = rows.Close()
	}()
	rList := make([]*object.HeartbeatMonitorData, 0)
	var clientId, coAb, coUserAb, svVer, clientVersion string
	var coId int
	var heartbeat time.Time
	outTime := time.Now().Add(-goToolCommon.GetDurationBySecond(global.SysConfig.SSConfig.HeartBeatForbidden))
	outTimeStr := goToolCommon.GetDateTimeStr(outTime)
	for rows.Next() {
		err := rows.Scan(&clientId, &coId, &coAb, &coUserAb, &svVer, &heartbeat, &clientVersion)
		if err != nil {
			errMsg := fmt.Sprintf("%s read data err: %s", "GetHeartbeatMonitorData", err.Error())
			log.Error(errMsg)
			return nil, errors.New(errMsg)
		}
		isOffLine := ""
		if goToolCommon.GetDateTimeStr(heartbeat) <= outTimeStr {
			isOffLine = "true"
		} else {
			isOffLine = "false"
		}
		rList = append(rList, &object.HeartbeatMonitorData{
			ClientId:      clientId,
			CoId:          coId,
			CoAb:          coAb,
			CoUserAb:      coUserAb,
			SvVer:         svVer,
			HeartBeat:     goToolCommon.GetDateTimeStrWithMillisecond(heartbeat),
			ClientVersion: clientVersion,
			IsOffLine:     isOffLine,
		})
	}

	if rows.Err() != nil {
		errMsg := fmt.Sprintf("GetHeartbeatMonitorData %s err: %s", cType, rows.Err().Error())
		log.Error(errMsg)
		return nil, errors.New(errMsg)
	}
	return rList, nil
}

func (r *repLocal) DelHeartbeat(clientId string) error {
	log.Debug("del ID " + clientId)
	err := goToolMSSqlHelper.SetRowsBySQL(r.dbConfig, sqlDelHeartbeat, clientId)
	if err != nil {
		errMsg := fmt.Sprintf("DelHeartbeat err: %s", err.Error())
		log.Error(errMsg)
		return errors.New(errMsg)
	}
	return nil
}

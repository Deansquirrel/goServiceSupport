package object

import "time"

type Response struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

type VersionResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Version string `json:"version"`
}

type TypeResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Type    string `json:"type"`
}

type ClientInfoRequest struct {
	ClientId      string `json:"clientid"`
	ClientType    string `json:"clienttype"`
	ClientVersion string `json:"clientversion"`
	HostName      string `json:"hostname"`
	DbId          int    `json:"dbid"`
	DbName        string `json:"dbname"`
	InternetIP    string `json:"internetip"`
}

type SvrV3InfoRequest struct {
	ClientId   string `json:"clientid"`
	CoId       int    `json:"coid"`
	CoAb       string `json:"coab"`
	CoCode     string `json:"cocode"`
	CoUserAb   string `json:"couserab"`
	CoUserCode string `json:"cousercode"`
	CoFunc     int    `json:"cofunc"`
	SvName     string `json:"svname"`
	SvVer      string `json:"svver"`
	//2019-08-01T17:58:12+08:00
	SvDate time.Time `json:"svdate"`
}

type SvrZ5ZlVersionRequest struct {
	ClientId      string    `json:"clientid"`
	ObjectName    string    `json:"objectname"`
	ObjectType    string    `json:"objecttype"`
	ObjectVersion string    `json:"objectversion"`
	ObjectDate    time.Time `json:"objectdate"`
}

type SvrZ5ZlCompanyRequest struct {
	ClientId    string    `json:"clientid"`
	CoId        int       `json:"coid"`
	CoAb        string    `json:"coab"`
	CoCode      string    `json:"cocode"`
	CoType      int       `json:"cotype"`
	CoUserAb    string    `json:"couserab"`
	CoUserCode  string    `json:"cousercode"`
	CoAccCrDate time.Time `json:"coacccrdate"`
}

type HeartBeatRequest struct {
	ClientId        string    `json:"clientid"`
	HeartBeatClient time.Time `json:"heartbeatclient"`
}

type HeartBeatResponse struct {
	ErrCode         int    `json:"errcode"`
	ErrMsg          string `json:"errmsg"`
	IsForbidden     int    `json:"Isforbidden"`
	ForbiddenReason string `json:"forbiddenreason"`
}

type JobRecordRequest struct {
	JobId    string `json:"jobid"`
	ClientId string `json:"clientid"`
	JobKey   string `json:"jobkey"`
}

type JobErrRecordRequest struct {
	JobId     string    `json:"jobid"`
	ErrMsg    string    `json:"errmsg"`
	OccurTime time.Time `json:"occurtime"`
}

//===================================================================================
type WelcomeDataRequest struct {
	HeartbeatClientType string `json:"heartbeatClientType"`
}

type HeartbeatErrCount struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type WelcomeDataResponse struct {
	ErrCode       int                 `json:"errcode"`
	ErrMsg        string              `json:"errmsg"`
	HeartbeatData []HeartbeatErrCount `json:"heartbeatData"`
}

type HeartbeatMonitorDataRequest struct {
	Type string `json:"type"`
}

type HeartbeatMonitorData struct {
	ClientId      string `json:"clientId"`
	CoId          int    `json:"coId"`
	CoAb          string `json:"coAb"`
	CoUserAb      string `json:"coUserAb"`
	SvVer         string `json:"svVer"`
	HeartBeat     string `json:"heartbeat"`
	ClientVersion string `json:"clientVersion"`
	IsOffLine     string `json:"isOffLine"`
}

type HeartbeatMonitorDataResponse struct {
	ErrCode int                    `json:"errcode"`
	ErrMsg  string                 `json:"errmsg"`
	Data    []HeartbeatMonitorData `json:"data"`
}

type DelHeartbeatRequest struct {
	ClientId string `json:"clientId"`
}

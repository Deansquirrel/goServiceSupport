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

//type ClientIdResponse struct {
//	ErrCode int    `json:"errcode"`
//	ErrMsg  string `json:"errmsg"`
//	Id      string `json:"id"`
//}

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

type HeartBeatRequest struct {
	ClientId        string    `json:"clientid"`
	HeartBeatClient time.Time `json:"heartbeatclient"`
}

type JobRecordRequest struct {
	JobId    string `json:"jobid"`
	ClientId string `json:"clientid"`
	JobKey   string `json:"jobkey"`
}

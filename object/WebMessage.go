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

type ClientIdRequest struct {
	ClientType string `json:"clienttype"`
	HostName   string `json:"hostname"`
	DbId       int    `json:"dbid"`
	DbName     string `json:"dbname"`
}

type ClientIdResponse struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
	Id      string `json:"id"`
}

type ClientFlashInfoRequest struct {
	ClientId      string `json:"clientid"`
	ClientVersion string `json:"clientversion"`
	InternetIP    string `json:"internetip"`
}

type SvrV3InfoRequest struct {
	ClientId   string
	CoId       int
	CoAb       string
	CoCode     string
	CoUserAb   string
	CoUserCode string
	CoFunc     int
	SvName     string
	SvVer      string
	//2019-08-01T17:58:12+08:00
	SvDate time.Time
}

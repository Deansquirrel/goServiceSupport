package object

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

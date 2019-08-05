package object

import "time"

type ClientTypeInfo struct {
	ClientType  string
	IsSvrV3     int
	HasDb       int
	LastVersion string
}

type ClientInfo struct {
	ClientId      string
	ClientType    string
	ClientVersion string
	HostName      string
	DbId          int
	DbName        string
	InternetIP    string
	LastUpdate    time.Time
}

type SvrV3Info struct {
	ClientId   string
	CoId       int
	CoAb       string
	CoCode     string
	CoUserAb   string
	CoUserCode string
	CoFunc     int
	SvName     string
	SvVer      string
	SvDate     time.Time
	LastUpdate time.Time
}

type HeartBeat struct {
	ClientId        string
	HeartBeatClient time.Time
	HeartBeat       time.Time
}

type JobRecord struct {
	JobId     string
	ClientId  string
	JobKey    string
	StartTime time.Time
	EndTime   time.Time
}

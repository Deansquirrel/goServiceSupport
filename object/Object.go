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

type ClientControl struct {
	ClientId        string
	IsForbidden     int
	ForbiddenReason string
	LastUpdate      time.Time
}

type JobErrRecord struct {
	JobId           string
	ErrHh           int
	ErrMsg          string
	OccurClientTime time.Time
	OccurTime       time.Time
}

type SvrZ5ZlVersion struct {
	ClientId      string
	ObjectName    string
	ObjectType    string
	ObjectVersion string
	ObjectDate    time.Time
}

type SvrZ5ZlCompany struct {
	ClientId    string
	CoId        int
	CoAb        string
	CoCode      string
	CoType      int
	CoUserAb    string
	CoUserCode  string
	CoAccCrDate time.Time
}

func GetJobErrRecordByRequest(d *JobErrRecordRequest, l int) []*JobErrRecord {
	if d == nil {
		return nil
	}
	rTime := time.Now()
	rList := make([]*JobErrRecord, 0)
	errMsg := d.ErrMsg
	currHh := 0
	currMsg := ""
	for {
		currMsg = ""
		if len(errMsg) <= l {
			currMsg = errMsg
			errMsg = ""
		} else {
			currMsg = errMsg[:l]
			errMsg = errMsg[l:]
		}
		rList = append(rList, &JobErrRecord{
			JobId:           d.JobId,
			ErrHh:           currHh,
			ErrMsg:          currMsg,
			OccurClientTime: d.OccurTime,
			OccurTime:       rTime,
		})
		currHh = currHh + 1
		if errMsg == "" {
			break
		}
	}
	return rList
}

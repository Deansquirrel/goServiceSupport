package repository

import (
	"github.com/Deansquirrel/goServiceSupport/global"
	"github.com/Deansquirrel/goToolMSSql"
)

type common struct {
}

func NewCommon() *common {
	return &common{}
}

//获取本地库连接配置
func (c *common) GetLocalDbConfig() *goToolMSSql.MSSqlConfig {
	return &goToolMSSql.MSSqlConfig{
		Server: global.SysConfig.LocalDb.Server,
		Port:   global.SysConfig.LocalDb.Port,
		DbName: global.SysConfig.LocalDb.DbName,
		User:   global.SysConfig.LocalDb.User,
		Pwd:    global.SysConfig.LocalDb.Pwd,
	}
}

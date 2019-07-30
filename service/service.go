package service

import (
	"github.com/Deansquirrel/goServiceSupport/global"
	"github.com/Deansquirrel/goServiceSupport/webServer"
	log "github.com/Deansquirrel/goToolLog"
)

//启动服务内容
func StartService() error {
	log.Debug("Start Service")
	defer log.Debug("Start Service Complete")

	go func() {
		ws := webServer.NewWebServer(global.SysConfig.Iris.Port)
		ws.StartWebService()
	}()

	return nil
}

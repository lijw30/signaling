package framework

import (
	"fmt"
	"github.com/lijw30/signaling/src/glog"
)

func Init(confFile string) error {
	conf := loadConf(confFile)
	fmt.Printf("conf:%+v\n", conf)

	glog.SetLogDir(conf.Log.LogDir)
	glog.SetLogFileName(conf.Log.LogFile)
	glog.SetLogLevel(conf.Log.LogLevel)
	glog.SetLogToStderr(conf.Log.LogToStderr)
	return nil
}

package models
import "github.com/astaxie/beego/logs"

var logger *logs.BeeLogger

func init() {
	logger = logs.NewLogger(1000)
	logger.SetLogger("console", `{"level": 3}`)
}

func GetLogger() *logs.BeeLogger {
	logger = logs.NewLogger(1000)
	logger.SetLogger("console", `{"level": 3}`)
	return logger
}

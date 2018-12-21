package commend

import "flag"

type config struct {
	name  string
	value interface{}
}

var DbEnvConfig config
var AppEnvConfig config

func init() {
	//启动配置项
	//命令行参数配置
	var AppEnvName = flag.String("appEnvName", "dev", "this is set application run env")
	var dbEnvName = flag.String("dbEnvName", "dev", "this is set db env info")
	AppEnvConfig = config{
		name:  "appEnvName",
		value: AppEnvName,
	}
	DbEnvConfig = config{
		name:  "dbEncName",
		value: dbEnvName,
	}
	//解析cmd
	flag.Parse()
}

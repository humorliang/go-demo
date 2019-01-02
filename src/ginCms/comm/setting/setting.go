package setting

import (
	"time"
	"github.com/go-ini/ini"
	"log"
)

type App struct {
	ArticleListItemNum int
}
type Server struct {
	RunMode      string
	HttpPort     int
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}
type Database struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var AppSetting = &App{}
var ServerSetting = &Server{}
var DatabaseSetting = &Database{}

//读取配置信息
//pro 为生产环境
func SetUp(mode string) {
	if mode == "pro" {
		//读取配置信息
		cfg, err := ini.Load("conf/pro.ini")
		if err != nil {
			log.Fatal("读取配置文件错误")
		}
		//分区与结构体映射
		cfg.Section("app").MapTo(AppSetting)
		cfg.Section("server").MapTo(ServerSetting)
		cfg.Section("database").MapTo(DatabaseSetting)
		ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
		ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	} else {
		cfg, err := ini.Load("conf/dev.ini")
		if err != nil {
			log.Fatal("读取配置文件错误")
		}
		//分区与结构体映射
		cfg.Section("app").MapTo(AppSetting)
		cfg.Section("server").MapTo(ServerSetting)
		cfg.Section("database").MapTo(DatabaseSetting)

		ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
		ServerSetting.WriteTimeout = ServerSetting.WriteTimeout * time.Second
	}
	//fmt.Println(DatabaseSetting, ServerSetting)
}

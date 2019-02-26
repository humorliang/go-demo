package setting

import (
	"time"
	"github.com/go-ini/ini"
	"github.com/smallnest/rpcx/log"
)

//app
type App struct {
	RuntimePath string
	LogPath     string
	JwtKey      string
}

//server
type Server struct {
	RunMode      string
	HttpPort     string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	TokenTimeout time.Duration
}

//数据库
type DB struct {
	Type     string
	User     string
	Password string
	Host     string
	Name     string
}

var (
	AppCfg    = &App{}
	ServerCfg = &Server{}
	DbCfg     = &DB{}
)

func ConfigInit(mode *string) {
	var err error
	//配置文件
	var f *ini.File
	//传入环境
	if *mode == "pro" {
		f, err = ini.Load("conf/pro.ini")
		if err != nil {
			log.Fatal("pro config file load err:", err)
		}
	} else {
		f, err = ini.Load("conf/dev.ini")
		if err != nil {
			log.Fatal("dev config file load err:", err)
		}
	}
	//关系映射
	err = f.Section("app").MapTo(AppCfg)
	err = f.Section("server").MapTo(ServerCfg)
	err = f.Section("database").MapTo(DbCfg)
	if err != nil {
		log.Fatal("setting section to map err:", err)
	}
	ServerCfg.WriteTimeout = ServerCfg.WriteTimeout * time.Second
	ServerCfg.ReadTimeout = ServerCfg.ReadTimeout * time.Second
	ServerCfg.TokenTimeout = ServerCfg.TokenTimeout * time.Hour
}

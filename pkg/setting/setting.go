package setting

import (
	"log"
	"time"

	"gopkg.in/ini.v1"
)


var (
	// Cfg 配置文件
	Cfg *ini.File
	// RunMode 运行模式
	RunMode string
	// HTTPPort 端口号
	HTTPPort int
	// ReadTimeout 读取超时时间
	ReadTimeout time.Duration
	// WaitTimeout 等待超时时间
	WaitTimeout time.Duration
	// PageSize 分页大小
	PageSize int
	// JwtSecret jwt密匙
	JwtSecret string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}
	LoadBase()
	LoadServer()
	LoadApp()
}

// LoadBase 加载基础配置
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// LoadServer 加载服务配置
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}
	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIMEOUT").MustInt(60)) * time.Second
	WaitTimeout = time.Duration(sec.Key("Wait_TIMEOUT").MustInt(60)) * time.Second
}

// LoadApp 加载App配置
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
	JwtSecret = sec.Key("JWT_SECRET").MustString("56354673456")
}
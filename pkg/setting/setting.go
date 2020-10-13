/** Created By wene<354007048@qq.com> . Date at 2020/6/2 */
package setting

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

type App struct {
	PageSize int
	RuntimeRootPath string

	Pbkdf2Secret string

	LogSavePath string
	LogSaveName string
	LogFileExt string
	TimeFormat string
}

var AppSetting = &App{}

type Jwt struct {
	JwtSecret string
	JwtRefreshSecret string
	JwtExpireTime time.Duration
	JwtRefreshExpireTime time.Duration
}

var JwtSetting = &Jwt{}

type File struct {
	FileDomain string
	FileSavePath string
	FileUploader string
	FileMultipartMaxSize int64
	FileMaxSize int
	FileNum int
	FileExclude []string
	FileInclude []string
}

var FileSetting = &File{}

type Server struct {
	RunMode string
	HttpPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
}

var ServerSetting = &Server{}

type Database struct {
	Type string
	User string
	Password string
	Host string
	Name string
	TablePrefix string
}

var DatabaseSetting = &Database{}

type Lin struct {
	GroupLevelRoot int
	GroupLevelGuest int
}

var LinSetting = &Lin{}

func Setup() {
	Cfg, err := ini.Load("internal/config/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'internal/config/app.ini': %v", err)
	}

	err = Cfg.Section("app").MapTo(AppSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo AppSetting err: %v", err)
	}

	err = Cfg.Section("file").MapTo(FileSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo FileSetting err: %v", err)
	}

	FileSetting.FileMultipartMaxSize = FileSetting.FileMultipartMaxSize * 1024 * 1024
	FileSetting.FileMaxSize = FileSetting.FileMaxSize * 1024 * 1024

	err = Cfg.Section("jwt").MapTo(JwtSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo JwtSetting err: %v", err)
	}

	JwtSetting.JwtExpireTime = JwtSetting.JwtExpireTime * time.Hour
	JwtSetting.JwtRefreshExpireTime = JwtSetting.JwtRefreshExpireTime * time.Hour

	err = Cfg.Section("server").MapTo(ServerSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo ServerSetting err: %v", err)
	}

	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second

	err = Cfg.Section("database").MapTo(DatabaseSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo DatabaseSetting err: %v", err)
	}

	err = Cfg.Section("lin").MapTo(LinSetting)
	if err != nil {
		log.Fatalf("Cfg.MapTo LinSetting err: %v", err)
	}
}

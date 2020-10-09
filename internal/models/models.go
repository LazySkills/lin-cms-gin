package models

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"lin-cms-gin/pkg/setting"
	stime "lin-cms-gin/pkg/time"
	"log"
	"os"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int            `gorm:"primary_key;column:id;type:int(11) unsigned auto_increment;not null;comment:'ID';" json:"id"`
	CreateTime stime.JSONTime `gorm:"column:create_time;type:datetime(3);default:CURRENT_TIMESTAMP(3);comment:'创建时间'"`
	UpdateTime stime.JSONTime `gorm:"column:update_time;type:datetime(3);default:CURRENT_TIMESTAMP(3) ON UPDATE CURRENT_TIMESTAMP(3);comment:'更新时间'"`
}

func Setup() {
	var err error

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second,   // 慢 SQL 阈值
			LogLevel:      logger.Info, // Log level
			Colorful:      false,         // 禁用彩色打印
		},
	)

	db, err = gorm.Open(mysql.New(mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
			setting.DatabaseSetting.User,
			setting.DatabaseSetting.Password,
			setting.DatabaseSetting.Host,
			setting.DatabaseSetting.Name),
		DefaultStringSize: 256, // string 类型字段的默认长度
		//DisableDatetimePrecision: true, // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		//DontSupportRenameIndex: true, // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		//DontSupportRenameColumn: true, // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		//SkipInitializeWithVersion: true, // 根据版本自动配置
	}), &gorm.Config{ //数据库配置，可以定义表，相关数据
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "",   // 表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: true, // 使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
		//NowFunc: func() time.Time {return time.Now().Local()},
		Logger: newLogger,
	})

	if err != nil {
		log.Println(err)
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(10) // SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(100) // SetMaxOpenConns 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour) // SetConnMaxLifetime 设置了连接可复用的最大时间
	AutoMigrateAll()
}

func AutoMigrateAll() {
	db.Set("gorm:table_options", "ENGINE=InnoDB")
	m := db.Migrator()

	if !m.HasTable(&Book{}) {
		m.AutoMigrate(&Book{})
		AddBook("庄稼汉","zds","asdasdkjhzcasnc","1.jpg")
	}

	if !m.HasTable(&LinFile{}) {
		m.AutoMigrate(&LinFile{})
	}

	if !m.HasTable(&LinGroup{}) {
		m.AutoMigrate(&LinGroup{})
		AddLinGroup("root","超级用户组",1)
		AddLinGroup("guest","游客组",2)
	}

	if !m.HasTable(&LinLog{}) {
		m.AutoMigrate(&LinLog{})
	}

	if !m.HasTable(&LinUser{}) {
		m.AutoMigrate(&LinUser{})
		AddLinUser("root","root","root@qq.com","1.jpg")
		AddLinUser("test","test","test@qq.com","2.jpg")
	}

	if !m.HasTable(&LinPermission{}) {
		m.AutoMigrate(&LinPermission{})
	}

	if !m.HasTable(&LinUserIdentity{}) {
		m.AutoMigrate(&LinUserIdentity{})
		AddIdentity(1,"USERNAME_PASSWORD","root","pbkdf2sha256:64000:18:24:n:yUnDokcNRbwILZllmUOItIyo9MnI00QW:6ZcPf+sfzyoygOU8h/GSoirF")
		AddIdentity(2,"USERNAME_PASSWORD","test","pbkdf2sha256:64000:18:24:n:yUnDokcNRbwILZllmUOItIyo9MnI00QW:6ZcPf+sfzyoygOU8h/GSoirF")
	}

	if !m.HasTable(&LinUserGroup{}) {
		m.AutoMigrate(&LinUserGroup{})
		AddLinUserGroup(1,1)
	}

	if !m.HasTable(&LinGroupPermission{}) {
		m.AutoMigrate(&LinGroupPermission{})
	}
}


func CloseDB() {
	sqlDB, _ := db.DB()
	defer sqlDB.Close()
}

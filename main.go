package main

import (
	"fmt"
	"lin-cms-gin/internal/models"
	"lin-cms-gin/internal/router"
	"lin-cms-gin/pkg/lin"
	"lin-cms-gin/pkg/logging"
	"lin-cms-gin/pkg/setting"
	"net/http"
)

func init() {
	setting.Setup()
	models.Setup()
	logging.Setup()
}

// dev win F:\tool\go\root_path\bin\fresh.exe 本地开发使用
// dev mac /Users/wangyu/go/bin/fresh 本地开发使用
func main() {
	if err := lin.InitTrans("zh"); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	router := router.InitRouter()

	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.ServerSetting.HttpPort),
		Handler:        router,
		ReadTimeout:    setting.ServerSetting.ReadTimeout,
		WriteTimeout:   setting.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}

	s.ListenAndServe()
}
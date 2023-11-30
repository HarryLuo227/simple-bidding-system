package main

import (
	"log"
	"net/http"
	"time"

	"github.com/HarryLuo227/simple-bidding-system/global"
	"github.com/HarryLuo227/simple-bidding-system/internal/model"
	"github.com/HarryLuo227/simple-bidding-system/internal/routers"
	"github.com/HarryLuo227/simple-bidding-system/pkg/setting"
	"github.com/gin-gonic/gin"
)

func init() {
	err := setupSetting()
	if err != nil {
		log.Fatalf("init.setupSetting err: %v", err)
	}

	err = setupDBEngine()
	if err != nil {
		log.Fatalf("init.setupDBEngine err: %v", err)
	}
}

func setupSetting() error {
	setting, err := setting.NewSetting("configs/")
	if err != nil {
		return err
	}

	err = setting.ReadSection("Server", &global.ServerSetting)
	if err != nil {
		return err
	}
	global.ServerSetting.ReadTimeout *= time.Second
	global.ServerSetting.WriteTimeout *= time.Second
	global.ServerSetting.TickerDuration *= time.Millisecond

	err = setting.ReadSection("Database", &global.DatabaseSetting)
	if err != nil {
		return err
	}

	return nil
}

func setupDBEngine() error {
	var err error
	global.DBEngine, err = model.NewDBEngine(global.DatabaseSetting)
	if err != nil {
		return err
	}
	return nil
}

func serveHome(c *gin.Context) {
	w := c.Writer
	r := c.Request
	http.ServeFile(w, r, "home.html")
}

func main() {
	router := routers.NewRouter()
	hub := newHub()
	go hub.run()
	router.GET("/", serveHome)
	router.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	})

	s := &http.Server{
		Addr:           ":" + global.ServerSetting.HttpPort,
		Handler:        router,
		ReadTimeout:    global.ServerSetting.ReadTimeout,
		WriteTimeout:   global.ServerSetting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

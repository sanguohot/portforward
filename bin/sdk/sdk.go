package sdk

import (
	"github.com/sanguohot/log/v2"
	"github.com/sanguohot/portforward/pkg/cfg"
	"github.com/sanguohot/portforward/pkg/forwarder"
	"os"
	"os/signal"
)

// for gomobile bind sdk
// 这里不能使用main方法
func Run() {
	cfg.LoadConfig()
	go forwarder.Serve()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit
	log.Sugar.Warnf("got os signal %v, forwarder exit", quit)
}

package forwarder

import (
	"github.com/1lann/udp-forward"
	"github.com/sanguohot/log/v2"
	"github.com/sanguohot/portforward/pkg/cfg"
	"github.com/sanguohot/portforward/pkg/common"
	"go.uber.org/zap"
	"net"
)

func Serve() {
	for _, v := range cfg.Config.Forwards {
		switch v.Network {
		case cfg.NetworkTcp:
			go serveTcp(v)
		case cfg.NetworkUdp:
			// 这里本来就是异步的
			serveUdp(v)
		}
	}
}

func serveUdp(forwardCfg cfg.ConfigForward) {
	_, err := forward.Forward(forwardCfg.SrcAddr, forwardCfg.DstAddr, forward.DefaultTimeout)
	if err != nil {
		log.Logger.Error(err.Error())
	}
}

func serveTcp(forward cfg.ConfigForward) {
	listener, err := net.Listen(forward.Network, forward.SrcAddr)
	if err != nil {
		log.Logger.Fatal(err.Error(), zap.String("network", forward.Network))
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Logger.Error(err.Error(), zap.String("network", forward.Network))
			continue
		}
		go handleConnection(conn, forward)
	}
}

func handleConnection(src net.Conn, forward cfg.ConfigForward) {
	dst, err := net.DialTimeout(forward.Network, forward.DstAddr, cfg.DialTimeout)
	if err != nil {
		log.Logger.Error(err.Error(), zap.String("network", forward.Network))
		return
	}
	common.Pipe(src, dst)
}

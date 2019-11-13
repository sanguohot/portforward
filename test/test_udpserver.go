package main

import (
	"github.com/sanguohot/log/v2"
	"net"
)

func handle(conn *net.UDPConn) {
	defer conn.Close()
	for {
		var buf [1024]byte
		n, addr, err := conn.ReadFromUDP(buf[0:])
		if err != nil {
			log.Logger.Error(err.Error())
			continue
		}
		log.Sugar.Infof("receive msg %s", string(buf[0:n]))
		_, err = conn.WriteToUDP([]byte("nice to see u"), addr)
		if err != nil {
			log.Logger.Error(err.Error())
			continue
		}
	}
}

func main() {
	addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:18000")
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	log.Sugar.Infof("listening udp on %v", addr)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		log.Logger.Error(err.Error())
		return
	}
	defer conn.Close()
	handle(conn)
}

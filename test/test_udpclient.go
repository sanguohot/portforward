package main

import (
	"github.com/sanguohot/log/v2"
	"net"
	"time"
)

func processUdp(dstAddr string) {
	conn, err := net.Dial("udp", dstAddr)
	defer conn.Close()
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	msg := "Hello world!"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	log.Sugar.Infof("send msg %s", msg)

	var res [1024]byte
	n, err := conn.Read(res[0:])
	if err != nil {
		log.Logger.Fatal(err.Error())
	}
	log.Sugar.Infof("receive msg %s", string(res[0:n]))
}

func main() {
	processUdp("localhost:18000")
	time.Sleep(time.Minute)
}

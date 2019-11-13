package common

import (
	"github.com/sanguohot/log/v2"
	"github.com/sanguohot/portforward/pkg/cfg"
	"io"
	"net"
	"sync"
	"time"
)

func printSocketErr(err error) {
	if err != nil {
		if err != io.EOF {
			log.Sugar.Debug(err.Error())
		}
	}
}

func Pipe(src, dst net.Conn) {
	defer func() {
		dst.Close()
		src.Close()
	}()
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		n, err := Copy(dst, src)
		printSocketErr(err)
		log.Sugar.Debugf("connection closed, src(%v) sent %d bytes to dst(%v)",
			src.RemoteAddr(), n, dst.RemoteAddr())
	}()
	go func() {
		defer wg.Done()
		n, err := Copy(src, dst)
		printSocketErr(err)
		log.Sugar.Debugf("connection closed, src(%v) received %d bytes from dst(%v)",
			src.RemoteAddr(), n, dst.RemoteAddr())
	}()
	wg.Wait()
}

// io.ReadWriteCloser
func Copy(dst, src net.Conn) (int64, error) {
	var (
		size int64
		err  error
		n    int
	)
	for {
		buf := make([]byte, cfg.BufferSize)
		src.SetReadDeadline(time.Now().Add(cfg.ReadTimeout))
		n, err = src.Read(buf)
		// 如果err==io.EOF，可能是读完了，但是有数据需要处理
		if err != nil {
			if err == io.EOF {
				if n == 0 {
					break
				}
			} else {
				break
			}
		}
		n, err = dst.Write(buf[:n])
		if err != nil {
			break
		}
		size += int64(n)
	}
	return size, err
}

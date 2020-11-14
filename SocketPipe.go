package SocketPipe

import (
	"net"
	"time"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Pipe(sockOne net.Conn, sockTwo net.Conn) {
	for {
		go handleSocks(sockOne, sockTwo)
		go handleSocks(sockTwo, sockOne)
	}
}
func handleSocks(sockOne net.Conn, sockTwo net.Conn) {
	var buffer []byte
	_ = sockOne.SetReadDeadline(time.Now().Add(5))
	_ = sockTwo.SetWriteDeadline(time.Now().Add(5))
	_, _ = sockOne.Read(buffer)
	_, _ = sockTwo.Write(buffer)
}

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
	err := sockOne.SetReadDeadline(time.Now().Add(5))
	check(err)
	err = sockTwo.SetWriteDeadline(time.Now().Add(5))
	check(err)
	num, err := sockOne.Read(buffer)
	check(err)
	if num > 0 {
		_, err := sockTwo.Write(buffer)
		check(err)
	}

}

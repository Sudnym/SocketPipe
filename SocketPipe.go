package SocketPipe

import (
	"bufio"
	"io"
	"net"
)

func Pipe(sockOne net.Conn, sockTwo net.Conn) {
	go handleSocks(sockOne, sockTwo)
	handleSocks(sockTwo, sockOne)
}
func handleSocks(sockOne net.Conn, sockTwo net.Conn) {
	var buffer []byte
	ch := make(chan []byte)
	for true {
		buffer, isClosed := read(sockOne, buffer, ch)
		if isClosed {
			sockTwo.Write([]byte("Other Client Closed\t"))
			return
		}
		_, err := sockTwo.Write(buffer)
		if err == io.EOF {
			return
			sockOne.Write([]byte("Other Client Closed\t"))
		}
	}
}
func read(sock net.Conn, buffer []byte, ch chan []byte) ([]byte, bool) {
	reader := bufio.NewReader(sock)
	isClosed := false
	buffer, err := reader.ReadBytes('\t')
	if err == io.EOF {
		isClosed = true
	}
	return buffer, isClosed
}

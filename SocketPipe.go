package SocketPipe

import (
	"bufio"
	"io"
	"net"
)

func Pipe(sockOne net.Conn, sockTwo net.Conn, delimiter byte) {
	go handleSocks(sockOne, sockTwo, delimiter)
	handleSocks(sockTwo, sockOne, delimiter)
}
func handleSocks(sockOne net.Conn, sockTwo net.Conn, delimiter byte) {
	var buffer []byte
	for true {
		buffer, isClosed := read(sockOne, buffer, delimiter)
		if isClosed {
			sockTwo.Write(append([]byte("Other Client Closed"), delimiter))
			return
		}
		_, err := sockTwo.Write(buffer)
		if err == io.EOF {
			sockOne.Write(append([]byte("Other Client Closed"), delimiter))
			return
		}
	}
}
func read(sock net.Conn, buffer []byte, delimiter byte) ([]byte, bool) {
	reader := bufio.NewReader(sock)
	isClosed := false
	buffer, err := reader.ReadBytes(delimiter)
	if err == io.EOF {
		isClosed = true
	}
	return buffer, isClosed
}

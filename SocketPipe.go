package SocketPipe

import (
	"fmt"
	"net"
	"sync"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func Pipe(sockOne net.Conn, sockTwo net.Conn) {
	fmt.Println(sockOne.RemoteAddr())
	fmt.Println(sockTwo.RemoteAddr())
	go handleSocks(sockOne, sockTwo)
	go handleSocks(sockTwo, sockOne)
}
func handleSocks(sockOne net.Conn, sockTwo net.Conn) {
	var buffer []byte
	var wg sync.WaitGroup
	for {
		wg.Add(1)
		go read(&wg, sockOne, buffer)
		wg.Wait()
		sockTwo.Write(buffer)
	}
}
func read(wg *sync.WaitGroup, sock net.Conn, buffer []byte) {
	sock.Read(buffer)
	wg.Done()
}

package SocketPipe

import (
	"fmt"
	"io/ioutil"
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
	ch := make(chan []byte)
	for {
		wg.Add(1)
		go read(&wg, sockOne, buffer, ch)
		wg.Wait()
		buffer = <-ch
		sockTwo.Write(buffer)
	}
}
func read(wg *sync.WaitGroup, sock net.Conn, buffer []byte, ch chan []byte) {
	defer wg.Done()
	buffer, err := ioutil.ReadAll(sock)
	ch <- buffer
	check(err)
}

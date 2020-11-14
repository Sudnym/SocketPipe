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
	handleSocks(sockTwo, sockOne)
}
func handleSocks(sockOne net.Conn, sockTwo net.Conn) {
	var buffer []byte
	var wg sync.WaitGroup
	ch := make(chan []byte)
	for true {
		wg.Add(1)
		go read(&wg, sockOne, buffer, ch)
		wg.Wait()
		buffer = <-ch
		sockTwo.Write(buffer)
	}
}
func read(wg *sync.WaitGroup, sock net.Conn, buffer []byte, ch chan []byte) {
	defer wg.Done()
	var buff []byte
	for string(buff) == string(buffer) {
		buffer, _ = ioutil.ReadAll(sock)
	}
	fmt.Println(buffer)
	ch <- buffer
}

package SocketPipe

import (
	"net"
	"time"
)

func Pipe(sockOne net.Conn, sockTwo net.Conn){
	sockOne.SetDeadline(time.Now().Add(0.1))
	sockTwo.SetDeadline(time.Now().Add(0.1))
	channelTo := make(chan []byte)
	channelFrom := make(chan []byte)
	var bufferTo []byte
	var bufferFrom []byte
	go sockTwo.Read(bufferTo)
	go func() { channelTo <- bufferTo }()
	go sockOne.Write(<- channelTo)
	go sockOne.Read(bufferFrom)
	go func() { channelFrom <- bufferFrom }()
	go sockTwo.Write(<- channelTo)
}
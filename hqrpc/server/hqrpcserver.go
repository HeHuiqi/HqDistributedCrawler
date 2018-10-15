package main

import (
	"net/rpc"
	"HqDistributedCrawler/hqrpc"
	"net"
	"log"
	"net/rpc/jsonrpc"
)

func main() {

	rpc.Register(hqrpc.DemoService{})
	listener,err := net.Listen("tcp",":1234")
	if err != nil {
		panic(err)
	}

	for {
		conn,err := listener.Accept()
		if err != nil {
			log.Printf("accept error: %v",err)
			continue
		}
		go jsonrpc.ServeConn(conn)
	}
}

//test
/*
telnet localhost 1234
{"method":"DemoService.Div","params":[{"A":3,"B":4}],"id":1}
{"method":"DemoService.Div","params":[{"A":3,"B":0}],"id":1234}


*/

package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type HelloWorld int

type Name struct {
	First string
	Last  string
}

func (h *HelloWorld) Hello(name *Name, reply *string) error {
	*reply = fmt.Sprintf("Hello %s %s!", name.First, name.Last)
	return nil
}

func main() {
	helloWorld := new(HelloWorld)
	rpc.Register(helloWorld)
	rpc.HandleHTTP()
	l, err := net.Listen("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(l, nil)
}

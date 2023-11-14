package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Name struct {
	First string
	Last  string
}

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dialing:", err)
	}
	name := Name{"Teddy", "Harbaugh"}
	var reply string
	err = client.Call("HelloWorld.Hello", name, &reply)
	if err != nil {
		log.Fatal("HelloWorld error:", err)
	}
	fmt.Println(reply)
}

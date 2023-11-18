package main

import (
	"fmt"
	"net/http"

	"github.com/quic-go/quic-go/http3"
	"github.com/quic-go/webtransport-go"
)

const (
	certFile = "certs/localhost.crt"
	keyFile  = "certs/localhost.key"
)

func main() {
	go frontendServer()
	webtransportServer()
}

func webtransportServer() {
	var webtrans = webtransport.Server{
		H3: http3.Server{Addr: "wsl.jph2.tech:8443"},
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}

	http.HandleFunc("/webtransport", func(w http.ResponseWriter, r *http.Request) {
		session, err := webtrans.Upgrade(w, r)
		if err != nil {
			fmt.Println("upgrading failed:", err)
			return
		}
		stream, err := session.OpenStream()
		if err != nil {
			fmt.Println("opening stream:", err)
			return
		}
		defer stream.Close()
		fmt.Fprintf(stream, "Do you read me?")
		var message = make([]byte, 80)
		for {
			n, _ := stream.Read(message)
			stream.Write(message[:n])
		}
	})

	err := webtrans.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		fmt.Println(err)
	}
	return
}

func frontendServer() {
	http.Handle("/", http.FileServer(http.Dir("frontend")))
	err := http.ListenAndServeTLS("wsl.jph2.tech:443", certFile, keyFile, nil)
	if err != nil {
		fmt.Println(err)
	}
}

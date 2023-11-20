package main

import (
	"io"
	"log"
	"net/http"

	"github.com/quic-go/quic-go/http3"
	"github.com/quic-go/webtransport-go"
)

const (
	certFile = "certs/jph2.tech.crt"
	keyFile  = "certs/jph2.tech.key"
)

var webtransportServer = webtransport.Server{
	H3: http3.Server{Addr: ":3122"},
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var (
	public  = make(chan []byte)
	clients = make(map[chan<- []byte]bool)
)

func main() {
	go serveFrontend()
	serveWebtransport()
}

func serveWebtransport() {
	go broadcaster()
	http.HandleFunc("/chat", handleWTConn)
	err := webtransportServer.ListenAndServeTLS(certFile, keyFile)
	if err != nil {
		log.Println("webtransport:", err)
	}
}

func handleWTConn(w http.ResponseWriter, r *http.Request) {
	session, err := webtransportServer.Upgrade(w, r)
	if err != nil {
		log.Println("webtransport: upgrade:", err)
		return
	}
	log.Printf("webtransport: opened session to %s\n", session.RemoteAddr())

	stream, err := session.OpenStream()
	if err != nil {
		log.Println("webtransport:", err)
		return
	}
	log.Printf("webtransport: opened stream %d\n", stream.StreamID())

	outgoing := make(chan []byte)
	go sendMessages(stream, outgoing)
	outgoing <- []byte("You are connected to the server.")
	clients[outgoing] = true
	err = readMessages(stream, public, session.RemoteAddr().String())
	if err != nil {
		delete(clients, outgoing)
		close(outgoing)
		stream.CancelWrite(0)
		log.Printf("webtransport: stream %d closed: %s\n", stream.StreamID(), err)
	}
}

func readMessages(stream webtransport.Stream, public chan<- []byte, name string) error {
	var message = make([]byte, 80)
	for {
		n, err := stream.Read(message)
		if err != nil && err != io.EOF {
			return err
		}
		public <- []byte(name + ": " + string(message[:n]))
		if err == io.EOF {
			return err
		}
	}
}

func sendMessages(stream webtransport.Stream, outgoing <-chan []byte) {
	for message := range outgoing {
		stream.Write(message)
	}
}

func broadcaster() {
	for message := range public {
		for client := range clients {
			client <- message
		}
	}
}

func serveFrontend() {
	http.Handle("/", http.FileServer(http.Dir("frontend")))
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Println("frontend:", err)
	}
}

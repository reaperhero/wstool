package main

import (
	"io"
	"log"
	"net"
	"net/http"
	"golang.org/x/net/websocket"
)

var  (
	tcpAddress = "testservicenats.myun.tv:4222"
	portString = ":90"
	binaryMode = false
)


func copyWorker(dst io.Writer, src io.Reader, doneCh chan<- bool) {
	io.Copy(dst, src)
	doneCh <- true
}

func relayHandler(ws *websocket.Conn) {
	conn, err := net.Dial("tcp", tcpAddress)
	if err != nil {
		log.Printf("[ERROR] %v \n", err)
		return
	}

	if binaryMode {
		ws.PayloadType = websocket.BinaryFrame
	}

	doneCh := make(chan bool)

	go copyWorker(conn, ws, doneCh)
	go copyWorker(ws, conn, doneCh)

	<-doneCh
	conn.Close()
	ws.Close()
	<-doneCh
}

func main() {
	log.Printf("[INFO] Listening on %s\n", portString)

	http.Handle("/", websocket.Handler(relayHandler))

	err := http.ListenAndServe(portString, nil)
	if err != nil {
		log.Fatal(err)
	}
}
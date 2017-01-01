package server

import (
	"fmt"
	"net"

	"log"

	"github.com/chaosvermittlung/schaschlik/config"
	"github.com/chaosvermittlung/schaschlik/writer"
)

var listener net.Listener
var size int

//Setup takes a confi.Config and Setups the TCP/UDP Server
func Setup(conf config.Config) {
	// Listen for incoming connections.
	var err error
	listener, err = net.Listen(conf.Type, ":"+conf.Port)
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}
	// Close the listener when the application closes.
	defer listener.Close()
	log.Println("Listening on :" + conf.Port)
	size = conf.Size
	go acceptConns()
}

func acceptConns() {
	for {
		// Listen for an incoming connection.
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Error accepting: ", err.Error())
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn, writer.Messages)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn, messages chan string) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, size)
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	} else {
		messages <- string(buf)
	}
	// Send a response back to person contacting us.
	conn.Write([]byte("Message received."))
	// Close the connection when you're done with it.
	conn.Close()
}

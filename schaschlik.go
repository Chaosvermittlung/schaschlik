package main

import (
	"flag"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	CONN_PORT = "2342"
	CONN_TYPE = "tcp"
)

var size = flag.Int("size", 1024, "Maxim payload per connection")
var path = flag.String("path", "test", "Path to write")

func main() {
	// Listen for incoming connections.
	flag.Parse()
	var messages = make(chan string)
	go printRequest(messages)
	l, err := net.Listen(CONN_TYPE, ":"+CONN_PORT)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	// Close the listener when the application closes.
	defer l.Close()
	fmt.Println("Listening on :" + CONN_PORT)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn, messages)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn, messages chan string) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, *size)
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

func printRequest(messages chan string) {
	for {
		m := <-messages
		fmt.Println("Message", messages)
		fmt.Println("path", *path)
		f, err := os.OpenFile(*path, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
		if err != nil {
			log.Fatal(err)
		}

		_, err = f.WriteString(m)
		if err != nil {
			log.Fatal(err)
		}
		f.Close()
	}
}

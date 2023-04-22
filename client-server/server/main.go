package main

import (
	"io"
	"net"
	"time"
)

func main() {
	// write server program to handle concurrent client
}

// handleConn - utility function
func handleConn(c net.Conn) {
	defer c.Close()
	for {
		_, err := io.WriteString(c, "response from server\n")
		if err != nil {
			return
		}
		time.Sleep(time.Second)
	}
}

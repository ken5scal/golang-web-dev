package main

import (
	"net"
	"log"
	"io"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	defer listener.Close()

	for {
		c, err := listener.Accept()
		if err != nil {
			log.Fatalln(err)
			continue
		}

		io.WriteString(c, "I see you connected.\n")
		c.Close()
	}
}

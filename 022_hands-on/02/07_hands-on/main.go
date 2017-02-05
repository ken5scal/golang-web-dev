package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
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

		go serve(c)
	}
}

func serve(c net.Conn) {
	defer c.Close()

	s := bufio.NewScanner(c)
	for s.Scan() {
		ln := s.Text()
		fmt.Println(ln)
		if ln == "" {
			// End of HTTP HEADER
			break
		}
	}
	fmt.Println("Code got here.")
	io.WriteString(c, "I see you connected.\n")
}

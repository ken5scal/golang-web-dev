package main

import (
	"net"
	"log"
	"io"
	"bufio"
	"fmt"
	"strings"
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
	var i int
	var method, url string
	for s.Scan() {
		ln := s.Text()

		if ln == "" {
			// End of HTTP HEADER
			break
		}

		if i == 0 {
			xs := strings.Fields(ln)
			method = xs[0]
			url = xs[1]
			fmt.Println("Method: ", method)
			fmt.Println("URL: ", url)
		}

		i++
	}
	fmt.Println("Code got here.")
	//io.WriteString(c, "I see you connected.\n")
	body := "Method is " + method + ". \nURL is " + url
	// I dont see difference between fmt.fprint and io.writestring
	fmt.Fprint(c, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(c, "Content-Length: %d\r\n", len(body))
	io.WriteString(c, "Content-Type: text/plain\r\n")
	io.WriteString(c, "\r\n")
	io.WriteString(c, body)
}

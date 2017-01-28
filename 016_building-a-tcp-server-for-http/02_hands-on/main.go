package main

import (
	"net"
	"log"
	"bufio"
	"fmt"
	"strings"
)

//create a server that returns the URL of the GET request
func main() {
	li, err := net.Listen("tcp", ":8080")
	checkError(err)
	defer li.Close()

	for {
		conn, err := li.Accept()
		if checkError(err) {
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()
	url := handleRequest(conn)     // read response
	handleResponse(conn, url) // write response
}
func handleResponse(conn net.Conn, url string) {

	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>` + url +  `Hello World</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func handleRequest(conn net.Conn) string {
	var url string
	i := 0
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			m := strings.Fields(ln)[0]
			url = strings.Fields(ln)[1]
			fmt.Println("***METHOD", m)
			fmt.Println("***URL", url)
			fmt.Println("***HTTP Version", strings.Fields(ln)[2])
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	return url
}



func checkError(e error) bool {
	if e != nil {
		log.Fatalln(e.Error())
		return true
	}
	return false
}

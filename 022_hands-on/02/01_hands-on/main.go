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

		s := bufio.NewScanner(c)
		for s.Scan() {
			fmt.Println(s.Text())
		}
		defer c.Close()

		fmt.Println("Code got here.")
		io.WriteString(c, "I see you connected.\n")

		c.Close()
	}
}

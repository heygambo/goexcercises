// Add code to WRITE to the connection.

package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Error starting the server", err)
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Fatalln("Error while accepting the connection", err)
		}
		go serve(conn)
	}
}

func serve(conn net.Conn) {
	defer conn.Close()
	s := bufio.NewScanner(conn)
	var i int
	for s.Scan() {
		ln := s.Text()

		if ln == "" {
			break
		}
		if i == 0 {
			ws := bufio.NewScanner(strings.NewReader(ln))
			ws.Split(bufio.ScanWords)
			var k int
			for ws.Scan() {
				w := ws.Text()
				switch k {
				case 0:
					fmt.Println("HTTP Method:", w)
				case 1:
					fmt.Println("HTTP URL:", w)
				}
				k++
			}
		} else {
			fmt.Println(ln)
		}
		i++
	}
	fmt.Println("Code got here.")
	io.WriteString(conn, "I see you connected\n")
}

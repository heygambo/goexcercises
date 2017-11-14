package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer li.Close()

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Fatalln(err.Error())
			continue
		}
		go handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	// read request
	request(conn)

	// write response
	// respond(conn)
}

func request(conn net.Conn) {
	i := 0
	scanner := bufio.NewScanner(conn)
	var url string
	for scanner.Scan() {
		ln := scanner.Text()
		fmt.Println(ln)
		if i == 0 {
			// request line
			url = strings.Fields(ln)[1]
			fmt.Println("***URL", url)
		}
		if ln == "" {
			// headers are done
			break
		}
		i++
	}
	switch url {
	case "/":
		respondHomepage(conn)
	case "/about":
		respondAbout(conn)
	default:
		respond404(conn)
	}
}

func respondHomepage(conn net.Conn) {
	fmt.Println("rendering homepage")
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Homepage</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func respondAbout(conn net.Conn) {
	fmt.Println("rendering about")
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>About us</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

func respond404(conn net.Conn) {
	fmt.Println("rendering 404")
	body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>404 Not found</strong></body></html>`

	fmt.Fprint(conn, "HTTP/1.1 404 Not found\r\n")
	fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
	fmt.Fprint(conn, "Content-Type: text/html\r\n")
	fmt.Fprint(conn, "\r\n")
	fmt.Fprint(conn, body)
}

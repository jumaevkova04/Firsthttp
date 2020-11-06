package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strconv"

	"github.com/jumaevkova04/http/pkg/server"
)

func main() {
	host := "0.0.0.0"
	port := "9999"

	if err := execute(host, port); err != nil {
		os.Exit(1)
	}
}

func execute(host string, port string) (err error) {
	srv := server.NewServer(net.JoinHostPort(host, port))

	srv.Register("/", func(conn net.Conn) {
		body := "Welcome to our web-site"

		_, err = conn.Write([]byte(fmt.Sprintf(
			"HTTP/1.1 200 OK\r\n"+
				"Contetnt-Length: "+strconv.Itoa(len(body))+"\r\n"+
				"Contetnt-Type: text/html\r\n"+
				"Connection: close\r\n") +
			"\r\n" +
			string(body),
		))
		if err != nil {
			log.Print(err)
		}
	})

	srv.Register("/about", func(conn net.Conn) {
		body := "About Golang Academy"

		_, err = conn.Write([]byte(fmt.Sprintf(
			"HTTP/1.1 200 OK\r\n" +
				"Contetnt-Length: " + strconv.Itoa(len(body)) + "\r\n" +
				"Contetnt-Type: text/html\r\n" +
				"Connection: close\r\n" +
				"\r\n" +
				string(body)),
		))
		if err != nil {
			log.Print(err)
		}
	})

	return srv.Start()
}

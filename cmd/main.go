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

	srv.Register("/payments", func(req *server.Request) {
		body := "About Golang Academy"

		id := req.QueryParams["id"]
		log.Print(id)

		_, err = req.Conn.Write([]byte(fmt.Sprintf(
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

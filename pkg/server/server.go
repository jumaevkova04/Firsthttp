package server

import (
	"bytes"
	"io"
	"log"
	"net"
	"strings"
	"sync"
)

// HandleFunc ...
type HandleFunc func(conn net.Conn)

// Server ...
type Server struct {
	addr     string
	mu       sync.RWMutex
	handlers map[string]HandleFunc
}

// NewServer ...
func NewServer(addr string) *Server {
	return &Server{addr: addr, handlers: make(map[string]HandleFunc)}
}

// Register ...
func (s *Server) Register(path string, handler HandleFunc) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.handlers[path] = handler
}

// Start ...
func (s *Server) Start() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Print(err)
		return err
	}

	defer func() {
		if cerr := listener.Close(); cerr != nil {
			err = cerr
			return
		}
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go s.handle(conn)

	}

}

func (s *Server) handle(conn net.Conn) {

	defer func() {
		if cerr := conn.Close(); cerr != nil {
			log.Print(cerr)
		}
	}()

	buf := make([]byte, 4096)

	n, err := conn.Read(buf)
	if err == io.EOF {
		log.Printf("%s", buf[:n])
	}
	if err != nil {
		log.Println(err)
		return
	}

	data := buf[:n]
	requestLineDelim := []byte{'\r', '\n'}
	requestLineEnd := bytes.Index(data, requestLineDelim)

	requestLine := string(data[:requestLineEnd])
	parts := strings.Split(requestLine, " ")

	path := parts[1]

	var handleFunc = func(conn net.Conn) {
		conn.Close()
	}

	if handler, ok := s.handlers[path]; ok {
		handleFunc = handler
		// return
	}

	handleFunc(conn)
}

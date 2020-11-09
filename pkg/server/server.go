package server

import (
	"bytes"
	"io"
	"log"
	"net"
	"net/url"
	"strings"
	"sync"
)

// Request ...
type Request struct {
	Conn        net.Conn
	QueryParams url.Values
	PathParams  map[string]string
	Headers     map[string]string
}

// HandleFunc ...
type HandleFunc func(req *Request)

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

	// p := strings.Split(path, "/")
	// // log.Print("first, p[2]: ", p[2])

	// s.handlers["/"+p[1]+"/"+"18"] = handler
	// log.Print("hahaha:  /" + p[1] + "/" + p[2])
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

	// log.Print("path: ", path)

	decoded, err := url.PathUnescape(path)
	if err != nil {
		log.Print(err)
		return
	}
	// log.Print(decoded)

	// p := strings.Split(decoded, "/")

	uri, err := url.ParseRequestURI(decoded)
	if err != nil {
		log.Print(err)
		return
	}
	// log.Print("uri.Path: ", uri.Path)
	// log.Print(uri.Query())

	var handleFunc = func(req *Request) {
		conn.Close()
	}

	if handler, ok := s.handlers[uri.Path]; ok {
		handleFunc = handler
	}

	var req Request
	req.Conn = conn
	req.QueryParams = uri.Query()

	headersLineDelim := []byte{'\r', '\n', '\r', '\n'}
	headersLineEnd := bytes.Index(data, headersLineDelim)

	headersLine := string(data[requestLineEnd:headersLineEnd])
	headers := strings.Split(headersLine, "\r\n")[1:]
	// headers = headers[1:]
	header := map[string]string{}
	for _, h := range headers {
		line := strings.Split(h, ": ")
		header[line[0]] = line[1]
	}
	req.Headers = header

	log.Print("header: ", header)

	// req.PathParams = map[string]string{"id": p[2]}

	handleFunc(&req)
}

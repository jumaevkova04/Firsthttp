package main

import (
	"net"
	"net/http"
	"os"

	"github.com/jumaevkova04/server/cmd/app"
	"github.com/jumaevkova04/server/pkg/banners"
)

//////////////////////////////////////////////////// Firts version ////////////////////////////////////////////////////

// type handler struct{}

// func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	_, err := w.Write([]byte("Hello world"))
// 	if err != nil {
// 		log.Print(err)
// 	}
// }

// func main() {
// 	host := "0.0.0.0"
// 	port := "9999"

// 	if err := execute(host, port); err != nil {
// 		os.Exit(1)
// 	}
// }

// func execute(host string, port string) (err error) {

// 	srv := http.Server{
// 		Addr:    net.JoinHostPort(host, port),
// 		Handler: &handler{},
// 	}

// 	return srv.ListenAndServe()
// }

//////////////////////////////////////////////////// Second version ////////////////////////////////////////////////////

// type handler struct {
// 	mu       *sync.RWMutex
// 	handlers map[string]http.HandlerFunc
// }

// func (h *handler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
// 	h.mu.RLock()
// 	handler, ok := h.handlers[request.URL.Path]
// 	h.mu.RUnlock()
// 	if !ok {
// 		http.Error(writer, http.StatusText(http.StatusNotFound), http.StatusNotFound)
// 		return
// 	}
// 	handler(writer, request)
// }

// func main() {
// 	host := "0.0.0.0"
// 	port := "9999"

// 	if err := execute(host, port); err != nil {
// 		os.Exit(1)
// 	}
// }

// func execute(host string, port string) (err error) {
// 	srv := http.Server{
// 		Addr:    net.JoinHostPort(host, port),
// 		Handler: &handler{},
// 	}
// 	return srv.ListenAndServe()
// }

//////////////////////////////////////////////////// Third version ////////////////////////////////////////////////////

// func main() {
// 	host := "0.0.0.0"
// 	port := "9999"

// 	if err := execute(host, port); err != nil {
// 		os.Exit(1)
// 	}

// }

// func execute(host string, port string) (err error) {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/banners.getAll", func(w http.ResponseWriter, r *http.Request) {
// 		_, err := w.Write([]byte("demo data"))
// 		if err != nil {
// 			log.Print(err)
// 		}
// 	})

// 	srv := http.Server{
// 		Addr:    net.JoinHostPort(host, port),
// 		Handler: mux,
// 	}

// 	return srv.ListenAndServe()
// }

//////////////////////////////////////////////////// Fourth version ////////////////////////////////////////////////////

func main() {
	host := "0.0.0.0"
	port := "9999"

	if err := execute(host, port); err != nil {
		os.Exit(1)
	}

}

func execute(host string, port string) (err error) {
	mux := http.NewServeMux()
	bannerSvc := banners.NewService()
	server := app.NewServer(mux, bannerSvc)
	server.Init()

	srv := &http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: server,
	}

	return srv.ListenAndServe()
}

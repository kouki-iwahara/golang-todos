package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Hello, World")

}

func main() {
	server := http.Server {
		Addr:              ":8080",
		// Handler:           nil,
		// TLSConfig:         &tls.Config{},
		// ReadTimeout:       0,
		// ReadHeaderTimeout: 0,
		// WriteTimeout:      0,
		// IdleTimeout:       0,
		// MaxHeaderBytes:    0,
		// TLSNextProto:      map[string]func( *http.Server,  *tls.Conn,  http.Handler){},
		// ConnState: func( net.Conn,  http.ConnState) {
		// },
		// ErrorLog: &log.Logger{},
		// BaseContext: func( net.Listener) context.Context {
		// },
		// ConnContext: func(ctx context.Context, c net.Conn) context.Context {
		// },
	}
	http.HandleFunc("/todos", handler)
	fmt.Println("server listening on port 8080")
	server.ListenAndServe()
}
package servers

import (
	"net/http"
	"time"
)

type HTTPServer struct {
	httpServer *http.Server
}

func NewHTTPServer(port string, handler http.Handler) *HTTPServer {
	return &HTTPServer{httpServer: &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		ReadTimeout:       10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
	}}
}

func (s *HTTPServer) Run() error {
	return s.httpServer.ListenAndServe()
}

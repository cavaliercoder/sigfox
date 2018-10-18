package sigfox

import (
	"net/http"
)

// A Server holds server data.
type Server struct{}

// ListenAndServe listens on the TCP network address.
func (s *Server) ListenAndServe(addr string) error {
	return http.ListenAndServe(addr, nil)
}

func (c *Server) ListenAndServerTLS(addr, certFile, keyFile string) error {
	return http.ListenAndServeTLS(addr, certFile, keyFile, nil)
}

func (c *Server) HandleUplink(pattern string, handler UplinkHandler) {
	http.Handle(pattern, handler)
}

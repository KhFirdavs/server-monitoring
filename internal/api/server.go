package api

import (
	"log"
	"net/http"
	"time"
)

type Server struct {
	server *http.Server
}

func (s *Server) ServerRun(handlers http.Handler, port string) error {
	s.server = &http.Server{
		Addr:         "localhost:" + port,
		Handler:      handlers,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	log.Println("server run success from port:", port)
	return s.server.ListenAndServe()
}

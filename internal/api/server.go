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
	log.Println("сервер запущен с портом:", port)
	return s.server.ListenAndServe()
}

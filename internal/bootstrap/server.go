package bootstrap

import (
	"fmt"
	"net/http"
)

type server struct {
	router http.Handler
	port   string
}

func (s *server) Run() error {
	fmt.Println("server running on port...")
	return http.ListenAndServe(s.port, s.router)
}

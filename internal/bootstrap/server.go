package bootstrap

import (
	"net/http"
)

type server struct {
	router http.Handler
	port   string
}

// func NewServer(router http.Handler, port string) *server {
// 	return &server{
// 		router: router,
// 		port:   fmt.Sprintf(":%s", port),
// 	}
// }

func (s *server) Run() error {
	return http.ListenAndServe(s.port, s.router)
}

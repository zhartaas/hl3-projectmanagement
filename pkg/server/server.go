package server

import (
	"fmt"
	"net/http"
)

type Server struct {
	http *http.Server
}

func New(handler http.Handler, port string) (r *Server, err error) {
	r = &Server{http: &http.Server{
		Handler: handler,
		Addr:    ":" + port,
	}}
	return
}

func (s *Server) Run() (err error) {
	go func() {
		if err = s.http.ListenAndServe(); err != nil {
			fmt.Println(err)
			return
		}
	}()

	return
}

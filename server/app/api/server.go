package main

import (
	"net"
	"net/http"

	v1 "github.com/ghostec/hall/server/app/api/v1"
	"github.com/gorilla/mux"
)

type Server struct {
	router   *mux.Router
	listener net.Listener
}

func NewServer() *Server {
	r := mux.NewRouter()

	v1.RegisterRoutes(r.PathPrefix("/v1").Subrouter())

	sv := Server{
		router: r,
	}

	return &sv
}

func (sv *Server) Listen(addr string) error {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	sv.listener = l
	return nil
}

func (sv *Server) Serve() error {
	httpSV := &http.Server{Handler: sv.router}
	return httpSV.Serve(sv.listener)
}

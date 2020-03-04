package core

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Server handle all server operations
type Server struct {
	Router *mux.Router
}

// ServerOpts enable set optional config to server
type ServerOpts struct {
	APIPrefix string
}

// Route describe a HTTP route
type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

// Bundle interface for app packages
type Bundle interface {
	GetRoutes() []Route
}

// NewServer load bundles, routes and return a instance of Server
func NewServer(bundles []Bundle, opts ServerOpts) *Server {
	r := mux.NewRouter()
	var s *mux.Router

	if opts.APIPrefix != "" {
		s = r.PathPrefix(opts.APIPrefix).Subrouter()
	} else {
		s = r
	}

	for _, b := range bundles {
		for _, route := range b.GetRoutes() {
			s.HandleFunc(route.Path, route.Handler).Methods(route.Method)
		}
	}

	http.Handle("/", r)

	return &Server{
		Router: r,
	}
}

// Start the server
func (srv *Server) Start(addr string) error {
	log.Printf("Listening on addr %s", addr)
	err := http.ListenAndServe(addr, nil)

	if err != nil {
		return err
	}

	return nil
}

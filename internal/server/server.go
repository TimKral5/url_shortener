// Package server handles all logic around the HTTP server and the
// coordination of all database connections.
package server

import (
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/timkral5/url_shortener/internal/auth"
	"github.com/timkral5/url_shortener/internal/cache"
	"github.com/timkral5/url_shortener/internal/database"
)

const requestTimeout time.Duration = 10 * time.Second
const maxHeaderSize = 4096

// Server is the wrapper for the main HTTP server.
type Server struct {
	server   *http.Server
	Database database.Connection
	Cache    cache.Connection
	Auth     auth.Connection
}

// NewServer constructs a new shortener API instance.
func NewServer() Server {
	api := Server{
		server:   nil,
		Database: nil,
		Cache:    nil,
		Auth:     nil,
	}

	return api
}

// CreateURLRoute creates a new shortened URL.
func (api *Server) CreateURLRoute(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// GetURLRoute fetches a full URL using its shortened ID.
func (api *Server) GetURLRoute(w http.ResponseWriter, r *http.Request) {
	id := strings.ToUpper(r.URL.Path[1:])

	_, err := w.Write([]byte(id))
	if err != nil {
		log.Println(err)
	}
}

// SetupRoutes constructs a serve mux and mounts all routes to it.
func (api *Server) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", api.CreateURLRoute)
	mux.HandleFunc("GET /", api.GetURLRoute)

	return mux
}

// Listen launches the HTTP server under the given address.
func (api *Server) Listen(addr string) bool {
	api.configureServer(addr)
	err := api.server.ListenAndServe()

	return err == nil
}

func (api *Server) configureServer(addr string) {
	api.server = &http.Server{
		ReadTimeout:                  requestTimeout,
		ReadHeaderTimeout:            requestTimeout,
		IdleTimeout:                  requestTimeout,
		Addr:                         addr,
		Handler:                      api.SetupRoutes(),
		DisableGeneralOptionsHandler: true,
		TLSConfig:                    nil,
		WriteTimeout:                 requestTimeout,
		MaxHeaderBytes:               maxHeaderSize,
		TLSNextProto:                 nil,
		ConnState:                    nil,
		ErrorLog:                     nil,
		BaseContext:                  nil,
		ConnContext:                  nil,
		HTTP2:                        nil,
		Protocols:                    nil,
	}
}

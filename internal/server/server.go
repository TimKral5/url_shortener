// Package server handles all logic around the HTTP server and the
// coordination of all database connections.
package server

import (
	"net/http"
	"strings"
	"time"

	"github.com/timkral5/url_shortener/internal/auth"
	"github.com/timkral5/url_shortener/internal/cache"
	"github.com/timkral5/url_shortener/internal/database"
	"github.com/timkral5/url_shortener/pkg/api"
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
	server := Server{
		server:   nil,
		Database: nil,
		Cache:    nil,
		Auth:     nil,
	}

	return server
}

// CreateURLRoute creates a new shortened URL.
func (server *Server) CreateURLRoute(writer http.ResponseWriter, _ *http.Request) {
	response := api.NewEmptyAddURLResponse()
	response.Hash = "Not Implemented"

	json, err := response.DumpJSON()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = writer.Write(json)
	if err != nil {
		return
	}
}

// GetURLRoute fetches a full URL using its shortened ID.
func (server *Server) GetURLRoute(writer http.ResponseWriter, request *http.Request) {
	id := strings.ToUpper(request.URL.Path[1:])

	response := api.NewEmptyGetURLResponse()
	response.URL = "Not Implemented (input: " + id + ")"

	json, err := response.DumpJSON()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = writer.Write(json)
	if err != nil {
		return
	}
}

// SetupRoutes constructs a serve mux and mounts all routes to it.
func (server *Server) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", server.CreateURLRoute)
	mux.HandleFunc("GET /", server.GetURLRoute)

	return mux
}

// Listen launches the HTTP server under the given address.
func (server *Server) Listen(addr string) bool {
	server.configureServer(addr)
	err := server.server.ListenAndServe()

	return err == nil
}

func (server *Server) configureServer(addr string) {
	server.server = &http.Server{
		ReadTimeout:                  requestTimeout,
		ReadHeaderTimeout:            requestTimeout,
		IdleTimeout:                  requestTimeout,
		Addr:                         addr,
		Handler:                      server.SetupRoutes(),
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

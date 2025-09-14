// Package server handles all logic around the HTTP server and the
// coordination of all database connections.
package server

import (
	"io"
	"net/http"
	"strings"
	"time"

	apidocs "github.com/timkral5/url_shortener/api"
	"github.com/timkral5/url_shortener/internal/auth"
	"github.com/timkral5/url_shortener/internal/cache"
	"github.com/timkral5/url_shortener/internal/database"
	"github.com/timkral5/url_shortener/internal/hash"
	"github.com/timkral5/url_shortener/internal/log"
	"github.com/timkral5/url_shortener/pkg/api"
)

const shortURLDefaultLength int = 10
const requestTimeout time.Duration = 10 * time.Second
const maxHeaderSize = 4096

// Server is the wrapper for the main HTTP server.
type Server struct {
	Database       database.Connection
	Cache          cache.Connection
	Auth           auth.Connection
	ShortURLLength int
	APIVersion     string
	server         *http.Server
}

// NewServer constructs a new shortener API instance.
func NewServer() *Server {
	server := &Server{
		Database:       nil,
		Cache:          nil,
		Auth:           nil,
		ShortURLLength: shortURLDefaultLength,
		APIVersion:     "NULL",
		server:         nil,
	}

	return server
}

// SetupRoutes constructs a serve mux and mounts all routes to it.
func (server *Server) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /v0/url", server.AddURLRoute)
	mux.HandleFunc("GET /v0/{file}", server.ServeDocsRoute)
	mux.HandleFunc("GET /{hash}", server.GetURLRoute)

	return mux
}

// AddURLRoute creates a new shortened URL.
func (server *Server) AddURLRoute(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("X-Api-Version", server.APIVersion)

	body, err := io.ReadAll(request.Body)
	if err != nil {
		log.Error("Read request body:", err)
		writer.WriteHeader(http.StatusBadRequest)

		return
	}

	hash, result := server.addURL(body)
	if !result {
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	server.handleAddURLResponse(writer, request, hash)
}

// GetURLRoute fetches a full URL using its shortened ID.
func (server *Server) GetURLRoute(writer http.ResponseWriter, request *http.Request) {
	id := request.PathValue("hash")

	writer.Header().Add("X-Api-Version", server.APIVersion)

	fullURL, result := server.getURL(id)
	if !result {
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	server.handleGetURLResponse(writer, request, fullURL)
}

func (server *Server) ServeDocsRoute(writer http.ResponseWriter, request *http.Request) {
	file := request.PathValue("file")

	writer.Header().Add("X-Api-Version", server.APIVersion)

	switch file {
	case "openapi.json":
		writer.Header().Add("Content-Type", "application/json")
		writer.Write([]byte(apidocs.JSONOpenAPISpecs))
	case "openapi.yaml":
		writer.Header().Add("Content-Type", "text/yaml")
		writer.Write([]byte(apidocs.YAMLOpenAPISpecs))
	default:
		writer.WriteHeader(http.StatusNotFound)
	}
}

// Listen launches the HTTP server under the given address.
func (server *Server) Listen(addr string) bool {
	server.configureServer(addr)
	err := server.server.ListenAndServe()
	log.Error(err)

	return err == nil
}

func (server *Server) trimHash(hash string) string {
	if len(hash) >= server.ShortURLLength {
		return hash[:server.ShortURLLength]
	}

	return hash
}

func (server *Server) addURL(body []byte) (string, bool) {
	requestData := api.NewEmptyAddURLRequest()

	err := requestData.LoadJSON(body)
	if err != nil {
		log.Error("Parsing request JSON:", err)

		return "", false
	}

	hash := strings.ToUpper(hash.GenerateSHA256Hex(requestData.URL))
	hash = server.trimHash(hash)

	err = server.Database.AddURL(hash, requestData.URL)
	if err != nil {
		log.Error("Adding URL to database:", err)

		return "", false
	}

	return hash, true
}

func (server *Server) handleAddURLResponse(writer http.ResponseWriter, _ *http.Request, hash string) {
	response := api.NewEmptyAddURLResponse()
	response.Hash = hash

	writer.Header().Add("Content-Type", "application/json")

	json, err := response.DumpJSON()
	if err != nil {
		log.Error("Constructing response:", err)
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	_, err = writer.Write(json)
	if err != nil {
		return
	}
}

func (server *Server) handleGetURLResponse(writer http.ResponseWriter, request *http.Request, url string) {
	acceptHeader := request.Header.Get("Accept")

	if strings.Contains(acceptHeader, "text/html") {
		writer.Header().Add("Content-Type", "text/html")
		writer.Header().Add("Location", url)
		writer.WriteHeader(http.StatusTemporaryRedirect)

		return
	}

	response := api.NewEmptyGetURLResponse()
	response.URL = url

	json, err := response.DumpJSON()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	writer.Header().Add("Content-Type", "application/json")

	_, err = writer.Write(json)
	if err != nil {
		return
	}
}

func (server *Server) getURL(hash string) (string, bool) {
	hash = server.trimHash(hash)

	fullURL, err := server.Cache.GetURL(hash)
	if err == nil {
		return fullURL, true
	}

	fullURL, err = server.Database.GetURL(hash)
	if err == nil {
		err = server.Cache.AddURL(hash, fullURL)
		if err != nil {
			log.Error("Failed to add URL to cache:", err)

			return "", false
		}

		return fullURL, true
	}

	log.Error("Failed to fetch URL from database:", err)

	return "", false
}

func (server *Server) configureServer(addr string) {
	server.server = &http.Server{
		ReadTimeout:                  requestTimeout,
		ReadHeaderTimeout:            requestTimeout,
		IdleTimeout:                  requestTimeout,
		Addr:                         addr,
		Handler:                      log.Middleware(server.SetupRoutes()),
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

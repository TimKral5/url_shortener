package api

import (
	"net/http"
	"strings"
)

type ShortenerApi struct {
	server http.Server
}

func NewShortenerApi() ShortenerApi {
	return ShortenerApi{}
}

func (self *ShortenerApi) CreateUrlRoute(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}

func (self *ShortenerApi) GetUrlRoute(w http.ResponseWriter, r *http.Request) {
	id := strings.ToUpper(r.URL.Path[1:])
	w.Write([]byte(id))
}

func (self *ShortenerApi) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /", self.CreateUrlRoute)
	mux.HandleFunc("GET /", self.GetUrlRoute)
	return mux
}

func (self *ShortenerApi) Listen(addr string) {
	self.server = http.Server{
		Addr: addr,
		Handler: self.SetupRoutes(),
	}
	self.server.ListenAndServe()
}

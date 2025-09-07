// Package main - This is the entrypoint of the main executable of
// this project.
package main

import (
	"github.com/timkral5/url_shortener/internal/cache"
	"github.com/timkral5/url_shortener/internal/database"
	"github.com/timkral5/url_shortener/internal/server"
)

func main() {
	server := server.NewServer()
	server.Cache = cache.NewFakeCacheConnection()
	server.Database = database.NewFakeDatabaseConnection()
	server.Listen(":3000")
}

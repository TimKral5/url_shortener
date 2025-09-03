// Package main - This is the entrypoint of the main executable of
// this project.
package main

import "github.com/timkral5/url_shortener/internal/server"

func main() {
	server := server.NewServer()
	server.Listen(":3005")
}

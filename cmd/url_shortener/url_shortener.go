// Package main - This is the entrypoint of the main executable of
// this project.
package main

import "github.com/timkral5/url_shortener/internal/api"

func main() {
	api := api.NewShortenerAPI()
	api.Listen(":3005")
}

package main

import "github.com/timkral5/url_shortener/internal/api"

func main() {
	api := api.NewShortenerApi()
	api.Listen(":3005")
}


package main

import (
	"net/http"
	"net/url"

	"github.com/taoqun8316/goblog/app/http/middlewares"
	"github.com/taoqun8316/goblog/bootstrap"
	"github.com/taoqun8316/goblog/pkg/logger"
)

type ArticlesFormData struct {
	Title, Body string
	URL         *url.URL
	Errors      map[string]string
}

type Article struct {
	Title, Body string
	ID          int64
}

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}

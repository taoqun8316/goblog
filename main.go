package main

import (
	"net/http"

	"github.com/taoqun8316/goblog/app/http/middlewares"
	"github.com/taoqun8316/goblog/bootstrap"
	"github.com/taoqun8316/goblog/pkg/logger"
)

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(":3000", middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}

package main

import (
	"net/http"

	"github.com/taoqun8316/goblog/app/http/middlewares"
	"github.com/taoqun8316/goblog/bootstrap"
	"github.com/taoqun8316/goblog/config"
	c "github.com/taoqun8316/goblog/pkg/config"
	"github.com/taoqun8316/goblog/pkg/logger"
)

func init() {
	config.Initialize()
}

func main() {
	bootstrap.SetupDB()
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}

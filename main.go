package main

import (
	"embed"
	"net/http"

	"github.com/taoqun8316/goblog/app/http/middlewares"
	"github.com/taoqun8316/goblog/bootstrap"
	"github.com/taoqun8316/goblog/config"
	c "github.com/taoqun8316/goblog/pkg/config"
	"github.com/taoqun8316/goblog/pkg/logger"
)

//go:embed resources/views/articles/*
//go:embed resources/views/auth/*
//go:embed resources/views/categories/*
//go:embed resources/views/layouts/*
var tplFS embed.FS

func init() {
	config.Initialize()
}

func main() {
	bootstrap.SetupDB()
	// 初始化模板
	bootstrap.SetupTemplate(tplFS)
	router := bootstrap.SetupRoute()

	err := http.ListenAndServe(c.GetString("app.port"), middlewares.RemoveTrailingSlash(router))
	logger.LogError(err)
}

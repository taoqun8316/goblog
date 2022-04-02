package bootstrap

import (
	"time"

	"github.com/gorilla/mux"
	"github.com/taoqun8316/goblog/pkg/model"
	"github.com/taoqun8316/goblog/pkg/route"
	"github.com/taoqun8316/goblog/routes"
)

// SetupRoute 路由初始化
func SetupRoute() *mux.Router {
	router := mux.NewRouter()
	routes.RegisterWebRoutes(router)
	route.SetRoute(router)
	return router
}

// SetupDB 初始化数据库和 ORM
func SetupDB() {
	db := model.ConnectDB()

	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxIdleTime(25)
	sqlDB.SetConnMaxIdleTime(5 * time.Minute)
}

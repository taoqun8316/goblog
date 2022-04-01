package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taoqun8316/goblog/app/http/controllers"
)

func RegisterWebRoutes(r *mux.Router) {

	pc := new(controllers.PagesController)
	// 静态页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/", pc.Home).Methods("GET").Name("home")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")

	// 文章相关页面
	ac := new(controllers.ArticlesController)
	r.HandleFunc("/articles/{id:[0-9]+}", ac.Show).Methods("GET").Name("articles.show")

}

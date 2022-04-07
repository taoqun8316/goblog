package routes

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taoqun8316/goblog/app/http/controllers"
	"github.com/taoqun8316/goblog/app/http/middlewares"
)

func RegisterWebRoutes(r *mux.Router) {
	r.Use(middlewares.StartSession)

	pc := new(controllers.PagesController)
	// 静态页面
	r.NotFoundHandler = http.HandlerFunc(pc.NotFound)
	r.HandleFunc("/", pc.Home).Methods("GET")
	r.HandleFunc("/about", pc.About).Methods("GET").Name("about")
	// 静态资源
	r.PathPrefix("/css/").Handler(http.FileServer(http.Dir("./public")))
	r.PathPrefix("/js/").Handler(http.FileServer(http.Dir("./public")))

	// 文章相关页面
	ac := new(controllers.ArticlesController)
	r.HandleFunc("/articles", middlewares.Auth(ac.Index)).Methods("GET").Name("home")
	r.HandleFunc("/articles/{id:[0-9]+}", middlewares.Auth(ac.Show)).Methods("GET").Name("articles.show")
	r.HandleFunc("/articles/create", middlewares.Auth(ac.Create)).Methods("GET").Name("articles.create")
	r.HandleFunc("/articles", middlewares.Auth(ac.Store)).Methods("POST").Name("articles.store")
	r.HandleFunc("/articles/{id:[0-9]+}/edit", middlewares.Auth(ac.Edit)).Methods("GET").Name("articles.edit")
	r.HandleFunc("/articles/{id:[0-9]+}", middlewares.Auth(ac.Update)).Methods("POST").Name("articles.update")
	r.HandleFunc("/articles/{id:[0-9]+}/delete", middlewares.Auth(ac.Delete)).Methods("POST").Name("articles.delete")

	// 文章分类
	cc := new(controllers.CategoriesController)
	r.HandleFunc("/categories/create", middlewares.Auth(cc.Create)).Methods("GET").Name("categories.create")
	r.HandleFunc("/categories", middlewares.Auth(cc.Store)).Methods("POST").Name("categories.store")
	r.HandleFunc("/categories/{id:[0-9]+}", cc.Show).Methods("GET").Name("categories.show")

	// 用户认证
	auc := new(controllers.AuthController)
	r.HandleFunc("/auth/register", middlewares.Guest(auc.Register)).Methods("GET").Name("auth.register")
	r.HandleFunc("/auth/do-register", middlewares.Guest(auc.DoRegister)).Methods("POST").Name("auth.doregister")
	r.HandleFunc("/auth/login", middlewares.Guest(auc.Login)).Methods("GET").Name("auth.login")
	r.HandleFunc("/auth/dologin", middlewares.Guest(auc.DoLogin)).Methods("POST").Name("auth.dologin")
	r.HandleFunc("/auth/logout", middlewares.Guest(auc.Logout)).Methods("POST").Name("auth.logout")

	// 用户相关
	uc := new(controllers.UserController)
	r.HandleFunc("/users/{id:[0-9]+}", uc.Show).Methods("GET").Name("users.show")

}

package controllers

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	"github.com/taoqun8316/goblog/pkg/logger"
	"github.com/taoqun8316/goblog/pkg/route"
	"github.com/taoqun8316/goblog/pkg/types"
)

type ArticlesController struct {
}

func Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	article, err := getArticleByID(id)

	if err != nil {
		if err == sql.ErrNoRows {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprint(w, "404 文章未找到")
		} else {
			logger.LogError(err)
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "500 服务器内部错误")
		}
	} else {
		tmpl, err := template.New("show.gohtml").Funcs(template.FuncMap{
			"RouteName2URL": route.Name2URL,
			"Int64ToString": types.Int64ToString,
		}).ParseFiles("resources/views/articles/show.gohtml")
		logger.LogError(err)

		err = tmpl.Execute(w, article)
		logger.LogError(err)
	}
}
package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/taoqun8316/goblog/app/models/article"
	"github.com/taoqun8316/goblog/app/policies"
	"github.com/taoqun8316/goblog/app/requests"
	"github.com/taoqun8316/goblog/pkg/auth"
	"github.com/taoqun8316/goblog/pkg/flash"
	"github.com/taoqun8316/goblog/pkg/logger"
	"github.com/taoqun8316/goblog/pkg/route"
	"github.com/taoqun8316/goblog/pkg/view"
)

type ArticlesController struct {
	BaseController
}

func (*ArticlesController) Index(w http.ResponseWriter, r *http.Request) {
	articles, err := article.GetAll()
	logger.LogError(err)

	if err != nil {
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	} else {
		view.Render(w, view.D{
			"Articles": articles,
		}, "articles.index", "articles._article_meta")
	}
}

func (ac *ArticlesController) Show(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	article, err := article.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		view.Render(w, view.D{
			"Article":          article,
			"CanModifyArticle": policies.CanModifyArticle(article),
		}, "articles.show", "articles._article_meta")
	}
}

func (*ArticlesController) Create(w http.ResponseWriter, r *http.Request) {
	view.Render(w, view.D{}, "articles.create", "articles._form_field")
}

func (*ArticlesController) Store(w http.ResponseWriter, r *http.Request) {
	currentUser := auth.User()
	_article := article.Article{
		Title:  r.PostFormValue("title"),
		Body:   r.PostFormValue("body"),
		UserID: currentUser.ID,
	}
	errors := requests.ValidateArticleForm(_article)

	// 检查是否有错误
	if len(errors) == 0 {
		_article.Create()

		if _article.ID > 0 {
			fmt.Fprint(w, "插入成功,ID为"+strconv.FormatInt(int64(_article.ID), 10))
		} else {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprint(w, "创建文章失败，请联系管理员")
		}
	} else {
		view.Render(w, view.D{
			"Title":  _article.Title,
			"Body":   _article.Body,
			"Errors": errors,
		}, "articles.create", "articles._form_field")
	}
}

func (ac *ArticlesController) Edit(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	article, err := article.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			view.Render(w, view.D{
				"Title":   article.Title,
				"Body":    article.Body,
				"Article": article,
				"Errors":  view.D{},
			}, "articles.edit", "articles._form_field")
		}
	}
}

func (ac *ArticlesController) Update(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_article, err := article.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			ac.ResponseForUnauthorized(w, r)
		} else {
			_article.Title = r.PostFormValue("title")
			_article.Body = r.PostFormValue("body")
			errors := requests.ValidateArticleForm(_article)

			if len(errors) == 0 {
				rowsAffected, err := _article.Update()

				if err != nil {
					logger.LogError(err)
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprint(w, "500 服务器内部错误")
					return
				}

				if rowsAffected > 0 {
					showURL := route.Name2URL("articles.show", "id", string(id))
					http.Redirect(w, r, showURL, http.StatusFound)
				} else {
					fmt.Fprint(w, "您没有做任何更改！")
				}
			} else {
				view.Render(w, view.D{
					"Title":  _article.Title,
					"Body":   _article.Body,
					"Errors": errors,
				}, "articles.edit", "articles._form_field")
			}
		}
	}
}

func (ac *ArticlesController) Delete(w http.ResponseWriter, r *http.Request) {
	id := route.GetRouteVariable("id", r)
	_article, err := article.Get(id)

	if err != nil {
		ac.ResponseForSQLError(w, err)
	} else {
		if !policies.CanModifyArticle(_article) {
			flash.Warning("您没有权限执行此操作！")
			http.Redirect(w, r, "/", http.StatusFound)
		} else {

			rowsAffected, err := _article.Delete()
			if err != nil {
				logger.LogError(err)
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprint(w, "500 服务器内部错误")
			} else {
				if rowsAffected > 0 {
					http.Redirect(w, r, route.Name2URL("articles.index"), http.StatusFound)
				} else {
					w.WriteHeader(http.StatusNotFound)
					fmt.Fprint(w, "404 文章未找到")
				}
			}
		}
	}
}

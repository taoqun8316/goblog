package controllers

import (
	"fmt"
	"net/http"
)

type PagesController struct {
}

func (p *PagesController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "欢迎来到go blog")
}

func (p *PagesController) About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "此博客是用以记录编程笔记，如您有反馈或建议，请联系 "+
		"<a href=\"mailto:summer@example.com\">summer@example.com</a>")
}

func (p *PagesController) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprint(w, "<h1>页面未找到</h1>")
}

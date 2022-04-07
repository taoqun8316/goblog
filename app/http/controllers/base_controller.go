package controllers

import (
	"fmt"
	"net/http"

	"github.com/taoqun8316/goblog/pkg/flash"
	"github.com/taoqun8316/goblog/pkg/logger"
	"gorm.io/gorm"
)

type BaseController struct {
}

func (bc BaseController) ResponseForSQLError(w http.ResponseWriter, err error) {
	if err == gorm.ErrRecordNotFound {
		// 3.1 数据未找到
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "404 文章未找到")
	} else {
		// 3.2 数据库错误
		logger.LogError(err)
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "500 服务器内部错误")
	}
}

func (bc BaseController) ResponseForUnauthorized(w http.ResponseWriter, r *http.Request) {
	flash.Warning("未授权操作！")
	http.Redirect(w, r, "/", http.StatusFound)
}

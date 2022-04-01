package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Name2URL(routName string, pairs ...string) string {
	var router *mux.Router
	url, err := router.Get(routName).URL(pairs...)
	if err != nil {
		return ""
	} else {
		return url.String()
	}
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

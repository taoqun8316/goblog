package route

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taoqun8316/goblog/pkg/logger"
)

var route *mux.Router

func SetRoute(r *mux.Router) {
	route = r
}
func Name2URL(routName string, pairs ...string) string {
	url, err := route.Get(routName).URL(pairs...)
	if err != nil {
		logger.LogError(err)
		return ""
	} else {
		return url.String()
	}
}

func GetRouteVariable(parameterName string, r *http.Request) string {
	vars := mux.Vars(r)
	return vars[parameterName]
}

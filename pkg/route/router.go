package route

import (
	"net/http"

	"github.com/gorilla/mux"
)

var Router mux.Router

func Initiallize() {
	Router = *mux.NewRouter()
}

func Name2URL(routName string, pairs ...string) string {
	url, err := Router.Get(routName).URL(pairs...)
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

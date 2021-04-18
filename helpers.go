package titan

import (
	"log"
	"net/url"

	"github.com/a-h/gemini"
	"github.com/a-h/gemini/mux"
)

func GetVar(r *gemini.Request, variable string) string {
	route, ok := mux.GetMatchedRoute(r.Context)
	if !ok {
		log.Fatal("couldn't parse route")
	}

	return route.PathVars[variable]
}

func GetQuery(r *gemini.Request) string {
	query, _ := url.QueryUnescape(r.URL.RawQuery)
	return query
}

package titan

import (
	"log"
	"net/url"

	"github.com/a-h/gemini"
	"github.com/a-h/gemini/mux"
)

func (server Booster) AddPage(path, file string, getData func(gemini.ResponseWriter, *gemini.Request) interface{}) {
	handler := func(w gemini.ResponseWriter, r *gemini.Request) {
		data := getData(w, r)
		TemplateIze(w, file, data)
	}
	server.Router.AddRoute(path, gemini.HandlerFunc(handler))
}

func (server Booster) AddAction(path string, handler func(gemini.ResponseWriter, *gemini.Request)) {
	server.Router.AddRoute(path, gemini.HandlerFunc(handler))
}

func (server Booster) AddInput(path, prompt string, handler func(gemini.ResponseWriter, *gemini.Request)) {
	internalHandler := func(w gemini.ResponseWriter, r *gemini.Request) {
		if GetQuery(r) == "" {
			w.SetHeader(gemini.CodeInput, prompt)
			return
		}
		handler(w, r)
	}
	server.Router.AddRoute(path, gemini.HandlerFunc(internalHandler))
}

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

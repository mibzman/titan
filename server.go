package titan

import (
	"context"
	"crypto/tls"
	"log"
	"net/url"

	"github.com/a-h/gemini"
	"github.com/a-h/gemini/mux"
	"gitlab.com/tslocum/gmitohtml/pkg/gmitohtml"
)

func GenerateServer() Server {
	return Server{mux.NewMux()}
}

type Server struct {
	Router *mux.Mux
}

func (server Server) Launch(domain string, cert tls.Certificate) error {
	gmitohtml.StartDaemon(domain+":80", domain, false)

	gem := gemini.NewDomainHandler(domain, cert, server.Router)

	err := gemini.ListenAndServe(context.Background(), ":1965", gem)
	if err != nil {
		log.Fatal("error:", err)
	}

	return err
}

func (server Server) AddPage(path, file string, getData func(gemini.ResponseWriter, *gemini.Request) interface{}) {
	handler := func(w gemini.ResponseWriter, r *gemini.Request) {
		data := getData(w, r)
		TemplateIze(w, file, data)
	}
	server.Router.AddRoute(path, gemini.HandlerFunc(handler))
}

func (server Server) AddAction(path string, handler func(gemini.ResponseWriter, *gemini.Request)) {
	server.Router.AddRoute(path, gemini.HandlerFunc(handler))
}

func (server Server) AddInput(path, prompt string, handler func(gemini.ResponseWriter, *gemini.Request)) {
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

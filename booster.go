package titan

import (
	"github.com/a-h/gemini"
)

func (booster Booster) AddPage(path, file string, getData func(gemini.ResponseWriter, *gemini.Request) interface{}) {
	handler := func(w gemini.ResponseWriter, r *gemini.Request) {
		data := getData(w, r)
		TemplateIze(w, file, data)
	}
	booster.Router.AddRoute(path, gemini.HandlerFunc(handler))
}

func (booster Booster) AddAction(path string, handler func(gemini.ResponseWriter, *gemini.Request)) {
	booster.Router.AddRoute(path, gemini.HandlerFunc(handler))
}

func (booster Booster) AddInput(path, prompt string, handler func(gemini.ResponseWriter, *gemini.Request)) {
	internalHandler := func(w gemini.ResponseWriter, r *gemini.Request) {
		if GetQuery(r) == "" {
			w.SetHeader(gemini.CodeInput, prompt)
			return
		}
		handler(w, r)
	}
	booster.Router.AddRoute(path, gemini.HandlerFunc(internalHandler))
}

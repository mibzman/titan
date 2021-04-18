package titan

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"

	"github.com/a-h/gemini"
	"github.com/a-h/gemini/mux"
	"github.com/mibzman/gmitohtml/pkg/gmitohtml"
)

type Booster struct {
	Router *mux.Mux
	LaunchConfig
}

type LaunchConfig struct {
	HTTPServer   bool
	HTTPPort     int
	GeminiServer bool
	GeminiPort   int
}

func Startup() Booster {
	return Booster{
		mux.NewMux(),
		LaunchConfig{true, 80, true, 1965},
	}
}

func (server Booster) Launch(domain string, cert tls.Certificate) error {
	if server.HTTPServer {
		gmitohtml.StartDaemon(domain+":"+fmt.Sprint(server.HTTPPort), domain, false)
	}

	if server.GeminiServer {
		gem := gemini.NewDomainHandler(domain, cert, server.Router)

		err := gemini.ListenAndServe(context.Background(), ":"+fmt.Sprint(server.GeminiPort), gem)
		if err != nil {
			log.Fatal("error:", err)
			return err
		}
	}

	return nil
}

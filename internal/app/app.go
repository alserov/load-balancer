package app

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type App struct {
	addr  string
	proxy *httputil.ReverseProxy
}

func New(addr string) *App {
	srvUrl, err := url.Parse(addr)
	if err != nil {
		log.Println(err)
	}

	return &App{
		addr:  addr,
		proxy: httputil.NewSingleHostReverseProxy(srvUrl),
	}
}

func (a *App) Address() string {
	return a.addr
}

func (a *App) IsAlive() bool {
	return true
}

func (a *App) Serve(w http.ResponseWriter, r *http.Request) {
	a.proxy.ServeHTTP(w, r)
}

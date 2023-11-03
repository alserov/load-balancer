package main

import (
	"github.com/alserov/load-balancer/internal/app"
	"github.com/alserov/load-balancer/internal/balancer"
	"log"
	"net/http"
)

const PORT = "3000"

func main() {
	servers := []balancer.Server{
		app.New("https://www.npmjs.com/"),
		app.New("https://github.com/alserov"),
	}

	lb := balancer.New(PORT, servers)

	handleRedirect := func(w http.ResponseWriter, r *http.Request) {
		lb.ServeProxy(w, r)
	}

	http.HandleFunc("/", handleRedirect)

	log.Println("serving requests at ", PORT)

	err := http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}

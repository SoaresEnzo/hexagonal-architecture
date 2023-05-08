package server

import (
	"github.com/codegangsta/negroni"
	_ "github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	_ "github.com/gorilla/mux"
	"github.com/soaresenzo/hexagonal-architecture-go/adapters/web/handler"
	"github.com/soaresenzo/hexagonal-architecture-go/application"
	"log"
	"net/http"
	"os"
	"time"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func NewWebserver() *WebServer {
	return &WebServer{}
}

func (w WebServer) Serve() {

	r := mux.NewRouter()
	n := negroni.New(negroni.NewLogger())

	handler.NewProductHandlers(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":8080",
		Handler:           http.DefaultServeMux,
		ErrorLog:          log.New(os.Stderr, "log:", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

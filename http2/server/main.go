package main

import (
	"io"
	"log"
	"net/http"
)

type Server struct {
}

func (s Server) ServeHTTP(writer http.ResponseWriter, request *http.Request) {

}

func main() {

	server := NewServer()
	log.Println("Starting server...at ", server.Addr)
	log.Fatalf("Error while serving server ::: %v", server.ListenAndServeTLS("server.crt", "server.key"))

}

func NewServer() *http.Server {
	initServerDefault()

	server := &http.Server{
		Addr: ":9000",
	}
	return server
}

func initServerDefault() {

	http.HandleFunc("/gorik", http1)
}

func http1(w http.ResponseWriter, r *http.Request) {
	bodyReq, _ := io.ReadAll(r.Body)
	log.Println("Gotta headers.... ", r)
	log.Println("Gotta headers.... ", r.Header)
	log.Println("Gotta body .... ", bodyReq)
	log.Println("-=-=-=-=-=-=--=-=-=-=-=-=- sending respponse -=-=-=-=-=-=--=-=-=-=-=-=-")

	w.Write([]byte("hello world"))

	log.Println("-=-=-=-=-=-=--=-=-=-=-=-=- respponse SENT -=-=-=-=-=-=--=-=-=-=-=-=-", bodyReq)
}

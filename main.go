package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/rpc"
	"github.com/gorilla/rpc/json"
)

type JSONServer struct{}

type Args struct {
	Name string
}

type Response struct {
	Result string
}

func (s *JSONServer) Say(r *http.Request, args *Args, result *Response) error {
	log.Println("chegou")
	result.Result = "Freeza! Por que você matou o Kuririn?"
	return nil
}

func main() {
	Server()
}

func Server() {
	log.Println("Iniciando POC")
	s := rpc.NewServer()

	s.RegisterCodec(json.NewCodec(), "application/json")
	s.RegisterCodec(json.NewCodec(), "application/json;charset=UTF-8")
	s.RegisterService(new(JSONServer), "")

	log.Println("Serviço registrado")

	router := mux.NewRouter()
	router.Handle("/poc", s)

	// http.Handle("/poc", s)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Println("error")
	}
}

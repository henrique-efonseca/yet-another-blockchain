package server

import (
	"github.com/henrique-efonseca/yet-another-blockchain-framework/internal/api"
	"github.com/henrique-efonseca/yet-another-blockchain-framework/internal/blockchain"
	"log"
	"net/http"
)

type Server struct {
	API *api.API
}

func NewServer(bc *blockchain.Blockchain) *Server {
	api := api.NewAPI(bc)
	return &Server{
		API: api,
	}
}

func (s *Server) Start(addr string) {
	log.Printf("Server listening on %s", addr)
	log.Fatal(http.ListenAndServe(addr, s.API.Router))
}

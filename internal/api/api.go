package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/henrique-efonseca/yet-another-blockchain-framework/internal/blockchain"
	"log"
	"net/http"
)

type API struct {
	Blockchain *blockchain.Blockchain
}

func NewAPI(bc *blockchain.Blockchain) *API {
	return &API{Blockchain: bc}
}

func (api *API) GetBlocks(w http.ResponseWriter, r *http.Request) {
	log.Println("GetBlocks endpoint hit")
	blocks, err := api.Blockchain.GetAllBlocks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blocks)
}

func (api *API) MineBlock(w http.ResponseWriter, r *http.Request) {
	log.Println("MineBlock endpoint hit")
	data := r.URL.Query().Get("data")
	if data == "" {
		http.Error(w, "Missing data parameter", http.StatusBadRequest)
		return
	}

	err := api.Blockchain.AddBlock(data, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	blocks, err := api.Blockchain.GetAllBlocks()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(blocks)
}

func (api *API) ValidateBlockchain(w http.ResponseWriter, r *http.Request) {
	log.Println("ValidateBlockchain endpoint hit")
	isValid := api.Blockchain.IsValid()
	result := map[string]bool{"valid": isValid}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func NewRouter(api *API) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/blocks", api.GetBlocks).Methods("GET")
	r.HandleFunc("/mine", api.MineBlock).Methods("POST")
	r.HandleFunc("/validate", api.ValidateBlockchain).Methods("GET")
	return r
}

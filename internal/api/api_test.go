package api

import (
	"github.com/henrique-efonseca/yet-another-blockchain-framework/internal/blockchain"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBlocks(t *testing.T) {
	bc := blockchain.NewBlockchain("testdb")
	defer bc.Close()
	srv := NewServer(bc)

	req, err := http.NewRequest("GET", "/blocks", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(srv.GetBlocks)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}

func TestMineBlock(t *testing.T) {
	bc := blockchain.NewBlockchain("testdb")
	defer bc.Close()
	srv := NewServer(bc)

	req, err := http.NewRequest("POST", "/mine?data=TestBlock", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(srv.MineBlock)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusCreated)
	}
}

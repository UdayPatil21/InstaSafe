package api

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleTransactions(t *testing.T) {

	server := httptest.NewServer(http.HandleFunc(HandleTransactions()))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			resp.StatusCode, http.StatusOK)
	}
}
func TestHandleTransactionsRR(t *testing.T) {
	rr:=httptest.NewRecorder()
	req,err:=http.NewRequest(http.MethodPost,"",nil)
	if err != nil {
		t.Error(err)
	}
	HandleTransactions(rr,req)
	if rr.Result().StatusCode != http.StatusOK {
		t.Errorf("expected 200 got %d ",rr.Result().StatusCode)
	}
	defere rr.Result().Body().Close()

}

func TestHandleStatistics(t *testing.T) {
	server := httptest.NewServer(http.HandleFunc(HandleStatistics()))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			resp.StatusCode, http.StatusOK)
	}
}

func TestSetLocationHandler(t *testing.T) {
	server := httptest.NewServer(http.HandleFunc(SetLocationHandler()))
	resp, err := http.Get(server.URL)
	if err != nil {
		t.Error(err)
	}
	if resp.StatusCode != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			resp.StatusCode, http.StatusCreated)
	}
}
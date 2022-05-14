package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"net/http"
	"strconv"
)

type Service struct {
	router *mux.Router
	quotes []string
}

func (s *Service) register() {
	s.router = mux.NewRouter()
	s.router.HandleFunc("/quotes", s.HandleNewQuote).Methods("POST")
	s.router.HandleFunc("/quotes", s.HandleGetQuotes)
}

func (s *Service) HandleGetQuotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(s.quotes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	return
}

func (s *Service) HandleNewQuote(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		http.Error(w, "request body can't be nil", http.StatusBadRequest)
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	quote := string(body)
	s.quotes = append(s.quotes, quote)

	return
}

func (s *Service) run(port int) {
	address := ":" + strconv.Itoa(port)
	err := http.ListenAndServe(address, s.router)
	if err == http.ErrServerClosed {
		err = nil
	}
	return
}

func Run(port int) {
	var s Service
	s.quotes = initialQuotes
	s.register()
	s.run(port)
}

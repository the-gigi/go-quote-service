package service

import (
	"encoding/json"
	"github.com/gorilla/mux"
	//"io/ioutil"
	"net/http"
	"strconv"
)

type Service struct {
	router *mux.Router
	//quoteStore om.QuoteStore
}

func (s *Service) register() {
	s.router.HandleFunc("/quotes", s.HandleNewQuote).Methods("POST")
	s.router.HandleFunc("/quotes", s.HandleGetQuotes)
}

func (s *Service) HandleGetQuotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//quotes, err := s.quoteStore.GetQuotes()
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//	return
	//}

	quotes := []string{
		"quote-1",
		"quote-2",
	}

	err := json.NewEncoder(w).Encode(quotes)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}

func (s *Service) HandleNewQuote(w http.ResponseWriter, r *http.Request) {
	//if r.Body == nil {
	//	http.Error(w, "request body can't be nil", http.StatusBadRequest)
	//	return
	//}
	//
	//body, err := ioutil.ReadAll(r.Body)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}
	//
	//quote := string(body)
	//err = s.quoteStore.AddQuote(quote)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusBadRequest)
	//	return
	//}

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

func Run(port int, connectionString string) (err error) {
	//quoteStore, err := quote_store.NewQuoteStore(connectionString)
	//if err != nil {
	//	return
	//}

	s := Service{
		router: mux.NewRouter(),
		//quoteStore: quoteStore,
	}
	s.register()
	s.run(port)
	return
}

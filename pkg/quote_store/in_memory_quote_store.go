package quote_store

import (
	om "github.com/the-gigi/go-quote-service/pkg/object_model"
)

type inMemoryQuoteStore struct {
	quotes []string
}

func (s *inMemoryQuoteStore) GetQuotes() (quotes []string, err error) {
	return s.quotes, nil
}

func (s *inMemoryQuoteStore) AddQuote(quote string) (err error) {
	s.quotes = append(s.quotes, quote)
	return
}

func NewInMemoryQuoteStore() (store om.QuoteStore) {
	return &inMemoryQuoteStore{}
}

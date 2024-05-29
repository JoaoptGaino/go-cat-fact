package main

import (
	"context"
	"encoding/json"
	"net/http"
)

type ApiServer struct {
	service Service
}

func NewApiServer(service Service) *ApiServer {
	return &ApiServer{
		service: service,
	}
}

func (s *ApiServer) Start(port string) error {
	http.HandleFunc("/cat-fact", s.handleGetCatFact)
	return http.ListenAndServe(port, nil)
}

func (s *ApiServer) handleGetCatFact(rw http.ResponseWriter, r *http.Request) {
	fact, err := s.service.GetCatFact(context.Background())

	if err != nil {
		writeJSON(rw, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	writeJSON(rw, http.StatusOK, fact)
}

func writeJSON(rw http.ResponseWriter, s int, v any) error {
	rw.WriteHeader(s)
	rw.Header().Add("Content-Type", "application/json")
	return json.NewEncoder(rw).Encode(v)
}

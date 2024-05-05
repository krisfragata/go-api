package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJSON(w http.ResponseWriter, status int, v any) error{
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func(http.ResponseWriter, *http.Request) error

type apiError struct {
	Error string
}

func (s *ApiServer) NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusNotFound, apiError{Error: "Page not found"})
}

func (s *ApiServer) GlobalErrorHandler(w http.ResponseWriter, r * http.Request, err error) {
	log.Println("Error:", err)
	WriteJSON(w, http.StatusInternalServerError, apiError{Error: "Internal server error"})
} 

func makeHTTPHandleFunc(f apiFunc, s *ApiServer) http.HandlerFunc{
	return  func(w http.ResponseWriter, r *http.Request){
		if err := f(w, r); err != nil{
			s.GlobalErrorHandler(w, r, err)
		}
	}
}


type ApiServer struct {
	listenAddress string
}

func NewAPIServer(listenAddress string) *ApiServer {
	return &ApiServer{
		listenAddress: listenAddress,
	}
}

func (s *ApiServer) Run(){
	router := mux.NewRouter()
	router.HandleFunc("/flow", makeHTTPHandleFunc(s.handleGetFlow, s) )
	router.NotFoundHandler = http.HandlerFunc(s.NotFoundHandler)

	log.Println("Server running on port:", s.listenAddress)
	http.ListenAndServe(s.listenAddress, router)
}


func (s *ApiServer) handleGetFlow(w http.ResponseWriter, r *http.Request) error{
	log.Println("get all flow data")
	return nil
}
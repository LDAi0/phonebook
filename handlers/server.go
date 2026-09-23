package handlers

import (
	"errors"
	"net/http"

	"github.com/gorilla/mux"
)

type HTTPserver struct{
	httphandlers *HTTPhandlers
}

func NewHTTPserver (httphandlers *HTTPhandlers) HTTPserver{
	return HTTPserver{
		httphandlers: httphandlers,
	}
}

func (s *HTTPserver) StartServer() error{
	router := mux.NewRouter()
	router.Path("/phonebook").Methods("Post").HandlerFunc(s.httphandlers.HandlerAddNumber)
	router.Path("/phonebook").Methods("Get").HandlerFunc(s.httphandlers.HandlerGetNumbers)
	router.Path("/phonebook/{title}").Methods("Get").HandlerFunc(s.httphandlers.HandlerGetNumber)
	if err:= http.ListenAndServe(":9091",router);err!=nil{
		if errors.Is(err,http.ErrServerClosed){
			return nil
		}
		return err
	}
	return nil
}
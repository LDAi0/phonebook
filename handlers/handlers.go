package handlers

import (
	"Phonebook/structs"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
)

type HTTPhandlers struct{
	book *structs.Book
}

func makeErrorDto(err error) ErrorDTO{
	return ErrorDTO{
			Message: err.Error(),
			Date: time.Now(),
		}
}

func NewHTTPhandlers(book *structs.Book) *HTTPhandlers{
	return &HTTPhandlers{
		book: book,
	}
}


func Write(status int, w http.ResponseWriter, b []byte){
	if _, err := w.Write(b); err!=nil{
		fmt.Println("failed to write http response: ", err)
		return
	} else {
		w.WriteHeader(status)
		return
	}
}

func (h *HTTPhandlers) HandlerAddNumber(w http.ResponseWriter, r *http.Request) {
	var numberdto NumberDTO
	if err := json.NewDecoder(r.Body).Decode(&numberdto); err != nil{
		errdto := makeErrorDto(err)
		http.Error(w,errdto.ToString(),http.StatusBadRequest)
		return 
	}
	
	if err:= numberdto.ValidateForCreate(); err!=nil{
		errdto := makeErrorDto(err)
		http.Error(w,errdto.ToString(),http.StatusBadRequest)
		return 
	}
	number:=structs.NewNumber(numberdto.PhoneNumber,numberdto.Name,numberdto.Group)
	if err:=h.book.AddNumber(number); err!=nil{
		errdto:=makeErrorDto(err)
		if errors.Is(err, structs.ErrNumberAlreadyExist){
			http.Error(w,errdto.ToString(),http.StatusConflict)
		} else {
			http.Error(w,errdto.ToString(),http.StatusInternalServerError)
		}
		return 
	}
	b, err := json.MarshalIndent(number,"","    ")
	if err!=nil{
		http.Error(w,"UnknownPanicErrorInMarshalToJsonErrorStruct",http.StatusInternalServerError)
		return
	}
	
	Write(http.StatusCreated,w,b)
}



func (h *HTTPhandlers) HandlerGetNumber(w http.ResponseWriter, r *http.Request) {
	PhoneNumber := mux.Vars(r)["title"]
	number, err :=h.book.GetNumber(PhoneNumber)
	if err!= nil{
		errdto:=makeErrorDto(err)
		if errors.Is(err,structs.ErrNumberNotFound){
			http.Error(w,errdto.ToString(),http.StatusBadRequest)
			return
		} else {
			http.Error(w,errdto.ToString(),http.StatusInternalServerError)
			return
		}
	}
	b,err:=json.MarshalIndent(number,"","    ")
	if err!=nil{
		http.Error(w,"UnknownPanicErrorInMarshalToJsonErrorStruct",http.StatusInternalServerError)
		return
	}


	Write(200,w,b)
}

func (h *HTTPhandlers) HandlerGetNumbers(w http.ResponseWriter, r *http.Request) {
	numbers:=h.book.GetNumbers()
	b,err:=json.MarshalIndent(numbers,"","    ")
	if err!=nil{
		http.Error(w,"UnknownPanicErrorInMarshalToJsonErrorStruct",http.StatusInternalServerError)
		return
	}
	Write(200,w,b)
}

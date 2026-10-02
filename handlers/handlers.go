package handlers

import (
	"Phonebook/database"
	"Phonebook/structs"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type HTTPhandlers struct{
	db *database.Db
}

func makeErrorDto(err error) ErrorDTO{
	return ErrorDTO{
			Message: err.Error(),
			Date: time.Now(),
		}
}

func NewHTTPhandlers(db *database.Db) *HTTPhandlers{
	return &HTTPhandlers{
		db: db,
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

func (h *HTTPhandlers) HandlerRegisterUser(w http.ResponseWriter, r *http.Request) {
	var userdto UserDTO
	if err := json.NewDecoder(r.Body).Decode(&userdto); err != nil{
		errdto := makeErrorDto(err)
		http.Error(w,errdto.ToString(),http.StatusBadRequest)
		return
	}

	if err := userdto.ValidateForRegister(); err!=nil{
		errdto:=makeErrorDto(err)
		http.Error(w,errdto.ToString(),http.StatusBadRequest)
		return
	}

	var userId int
	
	query := `INSERT INTO users (name,address,password)
	VALUES ($1,$2,$3) RETURNING id`
	user := structs.NewUser(userdto.Name,userdto.Password,userdto.Address)

	err := h.db.Conn.QueryRow(
		context.Background(),
		query,
		user.Name,
		user.Password,
		user.Address,
	).Scan(&userId)


	b, err := json.MarshalIndent(user,"","	")
	if err != nil{
		http.Error(w,"UnknownPanicErrorInMarshalToJsonErrorStruct",http.StatusInternalServerError)
		return
	}
	Write(201,w,b)

}

func (h *HTTPhandlers) HandlerLoginUser(w http.ResponseWriter, r *http.Request) {
	var userdto UserDTO
	if err := json.NewDecoder(r.Body).Decode(&userdto); err != nil{
		errdto := makeErrorDto(err)
		http.Error(w,errdto.ToString(),http.StatusBadRequest)
		return
	}
	if err := userdto.ValidateForRegister(); err!=nil{
		errdto:=makeErrorDto(err)
		http.Error(w,errdto.ToString(),http.StatusBadRequest)
		return
	}
	user := structs.NewUser(userdto.Name,userdto.Password,userdto.Address)
	query := `INSERT INTO users (name,address,password)
	VALUES ($1,$2,$3)`
	if _, err := h.db.Conn.Exec(context.Background(),query,user.Name,user.Address,user.Password); err!=nil{
		errdto:=makeErrorDto(err)
		http.Error(w,errdto.ToString(),http.StatusInternalServerError)
		return
	}
	b, err := json.MarshalIndent(user,"","	")
	if err != nil{
		http.Error(w,"UnknownPanicErrorInMarshalToJsonErrorStruct",http.StatusInternalServerError)
		return
	}
	Write(201,w,b)

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
	query:= `INSERT INTO numbers (phone_number,name,group_name,date_added)
	VALUES ($1,$2,$3,$4)`
	h.db.Conn.Exec(context.Background(),query,number.PhoneNumber,number.Name,number.Group,number.DateAdded)
	// if err:=h.book.AddNumber(number); err!=nil{
	// 	errdto:=makeErrorDto(err)
	// 	if errors.Is(err, structs.ErrNumberAlreadyExist){
	// 		http.Error(w,errdto.ToString(),http.StatusConflict)
	// 	} else {
	// 		http.Error(w,errdto.ToString(),http.StatusInternalServerError)
	// 	}
	// 	return 
	// }
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

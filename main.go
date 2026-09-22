package main

import (
	"Phonebook/handlers"
	"Phonebook/structs"
	"fmt"
)


func main(){
	book := structs.NewBook()
	httphandlers := handlers.NewHTTPhandlers(&book)
	httpserver := handlers.NewHTTPserver(httphandlers)

	if err := httpserver.StartServer(); err!=nil{
		fmt.Println("failed to start server: ", err)
	}
}
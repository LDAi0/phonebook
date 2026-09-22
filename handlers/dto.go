package handlers

import (
	"encoding/json"
	"errors"
	"time"
)

type NumberDTO struct {
	PhoneNumber string
	Name string
	Group string
}

func (nd *NumberDTO) Validate() error{
	if nd.PhoneNumber == ""{
		return errors.New("PhoneNumber is empty")
	}
	if nd.Name == ""{
		return errors.New("Name is empty")
	}
	return nil
}

type ErrorDTO struct{
	Message string
	Date time.Time
}

func (e *ErrorDTO) ToString() string{
	b,err:=json.MarshalIndent(e,"","	")
	if err!=nil{
		return "UnknownPanicErrorInMarshalToJsonErrorStruct"
	}
	return string(b)
}
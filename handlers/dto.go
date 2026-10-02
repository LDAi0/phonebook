package handlers

import (
	"encoding/json"
	"errors"
	"time"
)

type NumberDTO struct {
	Id int
	PhoneNumber string
	Name string
	Group string
}

type UserDTO struct{
	Id int
	Name string `json:"name"`
	Password string `json:"password"`
	Address string `json:"address"`
}

func (us *UserDTO) ValidateForRegister() error{
	if us.Name == ""{
		return errors.New("Name of User is empty")
	}
	if us.Password==""{
		return errors.New("Password of User is empty")
	}
	return nil
}



func (nd *NumberDTO) ValidateForCreate() error{
	if nd.PhoneNumber == ""{
		return errors.New("PhoneNumber is empty")
	}
	if nd.Name == ""{
		return errors.New("Name is empty")
	}
	return nil
}

func (nd *NumberDTO) ValidateForGet() error{
	if nd.PhoneNumber == ""{
		return errors.New("PhoneNumber is empty")
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
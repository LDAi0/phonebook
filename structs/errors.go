package structs

import "errors"

var ErrNumberAlreadyExist error = errors.New("number already exist")
var ErrNumberNotFound error = errors.New("number not found")

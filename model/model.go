package model

import "errors"

var (
	//ErrPersonCanNotBeNil la persona no puede ser nula
	ErrPersonCanNotBeNil = errors.New("La persona no puede ser nula")
	//IDPersonNotFount la persona no existe 
	IDPersonNotFount = errors.New("La persona no existe")
	
)

package handler

import "github.com/basic_API_with_GO_and_postgres/model"

type Storage interface {
	Create(*model.Person)error
	Update(int, *model.Person)error
	Delete(int)error
	GetByID(int)(model.Person, error)
	GetAll()(model.Persons, error)
}


package model

/**
	Creamos el modelo de personas
**/
type Community struct {
	Name string
}

//slice de comunidades
type Communities []Community

type Person struct {
	Name        string      `json:"name"`
	Age         uint8       `json:"age"`
	Communities Communities `json:"communities"`
}

type Persons []Person

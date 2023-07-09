package storage
/**
Esta estrcutura nos permite almacenar en memoria la información de las personas


**/
import (
	"fmt"

	"github.com/basic_API_with_GO_and_postgres/model"
)

type Memory struct {
	CurrenID int
	Persons  map[int]model.Person
}

//constructor of memory
func NewMemory()*Memory{
	person:= make(map[int]model.Person)

	return &Memory{
		CurrenID: 0,
		Persons: person,
	}
}


func( m *Memory)Create(person *model.Person) error{
	if person == nil{
		return model.ErrPersonCanNotBeNil
	}

	m.CurrenID++
	m.Persons[m.CurrenID] = *person

	return nil
}

 func (m *Memory)Update(ID int, person *model.Person)error{
	if person == nil{
		return model.ErrPersonCanNotBeNil
	}

	if _, ok := m.Persons[ID]; !ok {
		return fmt.Errorf(" ID %d: %w",ID, model.IDPersonNotFount)
	}

	m.Persons[ID] = *person

	return nil
 }

 func (m *Memory)Delete(ID int)error{
	if _,ok:= m.Persons[ID]; !ok{
		return fmt.Errorf(" ID %d: %w",ID, model.IDPersonNotFount)

	}
	delete(m.Persons, ID)

	return nil
 }

 func (m *Memory)GetByID(ID int)(model.Person, error){
	if person, ok := m.Persons[ID]; !ok{
		return person, fmt.Errorf(" ID %d: %w",ID, model.IDPersonNotFount)
	}else{
		return person, nil
	}

	
 }

 func (m *Memory)GetAll()(model.Persons, error){
	var result model.Persons

	for _, v := range m.Persons{
		result = append(result, v)
	}
	return result, nil
 }
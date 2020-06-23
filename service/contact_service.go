package service

import (
	"fmt"
	"log"
	"github.com/google/uuid"
	"../dto"
	"../model"
	"../repository"
)

func CreateContact(contact *model.Contact, repository repository.ContactRepository) dto.Response {
	uuidResult, err := uuid.NewRandom()

	if err != nil {
		log.Fatal(err)
	}

	contact.ID = uuidResult.String()

	operationResult := repository.Save(contact)

	if operationResult.Error != nil {
		return dto.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*model.Contact)

	return dto.Response{Success: true, Data : data}
}

func FindAllContacts(repository repository.ContactRepository) dto.Response{
	operationResult := repository.FindAll()

	if operationResult.Error != nil {
		return dto.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var datas = operationResult.Result.(*model.Contacts)

	return dto.Response{Success: true, Data: datas}
}

func FindOneContactById(id string, repository repository.ContactRepository) dto.Response{
	operationResult := repository.FindOneById(id)

	if operationResult.Error != nil{
		return dto.Response{Success: false, Message: operationResult.Error.Error()}
	}

	var data = operationResult.Result.(*model.Contact)

	return dto.Response{Success: true, Data: data}
}

func UpdateContactById(id string, contact *model.Contact, repository repository.ContactRepository) dto.Response {
	existingContactResponse := FindOneContactById(id, repository)

	fmt.Println(contact)

	if !existingContactResponse.Success {
		return existingContactResponse
	}

	existingContact := existingContactResponse.Data.(*model.Contact)

	fmt.Println(existingContact.Name + " old name")

	existingContact.Name = contact.Name
	existingContact.Email = contact.Email

	operationResult := repository.Save(existingContact)

	if operationResult.Error != nil {
		return dto.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dto.Response{Success: true, Data: operationResult.Result}
}

func DeleteContactById(id string, repository repository.ContactRepository) dto.Response{
	operationResult := repository.DeleteOneById(id)

	if operationResult.Error != nil {
		return dto.Response{Success: false, Message: operationResult.Error.Error()}
	}

	return dto.Response{Success: true}
}


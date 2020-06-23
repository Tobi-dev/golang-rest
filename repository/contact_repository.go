package repository

import (
	"../model"
	"github.com/jinzhu/gorm"
)

type ContactRepository struct {
	db *gorm.DB
}

func NewContactRepository(db *gorm.DB) *ContactRepository {
	return &ContactRepository{db : db}
}

func (r *ContactRepository) Save(contact *model.Contact) RepositoryResult {
	err := r.db.Save(contact).Error

	if err != nil {
		return RepositoryResult{Error : err}
	}

	return RepositoryResult{Result : contact}
}

func (r *ContactRepository) FindAll() RepositoryResult {
	var contacts model.Contacts

	err := r.db.Find(&contacts).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &contacts}
}

func (r *ContactRepository) FindOneById(id string) RepositoryResult {
	var contact model.Contact

	err := r.db.Where(&model.Contact{ID: id}).Take(&contact).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: &contact}
}

func (r *ContactRepository) DeleteOneById(id string) RepositoryResult {
	err := r.db.Delete(&model.Contact{ID: id}).Error

	if err != nil {
		return RepositoryResult{Error: err}
	}

	return RepositoryResult{Result: nil}
}
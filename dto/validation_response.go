package dto

type ValidationResponse struct {
	Success bool `json:"success"`
	Validations []Validation `json:"validations"`
}
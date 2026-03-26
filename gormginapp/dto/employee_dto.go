package dto

type CreateEmployeeRequest struct {
	Name string `json:"name" validate:"required,min=4"`
	City string `json:"city" validate:"required"`
}
type UpdateEmployeeRequest struct {
	Name string `json:"name"`
	City string `json:"city"`
}

package dtos

// DTOInterface defines a contract for Data Transfer Objects (DTOs) that
// require validation. Any struct implementing this interface must provide
// a Validate method that returns an error if the DTO is not valid.

type DTOInterface interface {
	Validate() error
}
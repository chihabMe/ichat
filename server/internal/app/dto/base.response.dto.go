package dto


type Status string
const (
	StatusError Status="error"
	StatusSuccess Status="success"
)

type BaseResponseDTO struct {
	Message string
	Status Status
	Errors error
	Data interface{}  
}
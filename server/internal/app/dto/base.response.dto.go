package dto


type Status string
const (
	StatusError Status="error"
	StatusSuccess Status="success"
)

type BaseResponseDTO struct {
	Message string `json:"message"`
	Status Status `json:"status"`
	Errors error `json:"errors"`
	Data interface{}  `json:"data"`
}
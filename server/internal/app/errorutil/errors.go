package errorutil

import "github.com/gofiber/fiber/v2"

type CustomError struct {
	Message string 
	Status int
}
type BodyValidationError struct {
	CustomError
	Errors interface{}
}

func (e *CustomError) Error()string{
	return e.Message
}


func NewCustomError(message string,status int)*CustomError{
	return &CustomError{Message: message,Status:status }
}
func NewValidationError(err interface{},message ...string)*BodyValidationError{
	errorMessage := "Invalid fields"
	if len(message)>0{
		errorMessage = message[0]
	}
	return &BodyValidationError{
		CustomError: CustomError{
			Message:errorMessage,
			Status: fiber.StatusBadRequest,
		},
		Errors: err,
	}
}

var (
	ErrFailedToParseData  = &CustomError{Message: "Failed to parse data",Status: fiber.StatusBadRequest}
	ErrInternalServerError = &CustomError{Message: "Internal server error",Status: fiber.StatusInternalServerError}
)

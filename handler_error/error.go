package handlererror
import (
	
)

type ResponseMessage struct {
	Message string
}

type ErrorMessage struct {
	Message string
	error
}

func NewResponseMessage(message string) ResponseMessage {
	return ResponseMessage{
		Message: message,
	}
}

func NewErrorMessage(message string) ErrorMessage {
	return ErrorMessage{
		Message: message,
	}
}

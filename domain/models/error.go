package models

type DomainError struct {
	Message string
}

func (e DomainError) Error() string {
	return e.Message
}

func NewOrderHadSubmittedError() error {
	msg_error := "order's id had submitted"
	return DomainError{
		Message: msg_error,
	}
}

func NewOrderCustomerFieldIsEmptyError() error {
	msg_error := "customer field is empty"
	return DomainError{
		Message: msg_error,
	}
}

func NewOrderOptionsFieldLessThanEqualToZeroError() error {
	msg_error := "options field less than equal to empty"
	return DomainError{
		Message: msg_error,
	}
}

func NewOrderIdNotFoundError() error {
	msg_error := "order's id not found"
	return DomainError{
		Message: msg_error,
	}
}

func NewUnableToCreateDuplicateOrder() error {
	msg_error := "unable to create duplicate order"
	return DomainError{
		Message: msg_error,
	}
}

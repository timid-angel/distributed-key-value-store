package domain

type DomainError struct {
	message string
}

func (err *DomainError) Error() string {
	return err.message
}

func NewDomainError(message string) DomainError {
	return DomainError{message: message}
}

type IService interface {
	HandleGet(key string) DomainError
	HandlePut(key string, value string) DomainError
	HandleDelete(key string) DomainError
	HandleList() DomainError
}

type IController interface {
	Get(key string) string
	Put(key string, value string) string
	Delete(key string) string
	List() string
}

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
	Get(key string) DomainError
	Put(key string, value string) DomainError
	Delete(key string) DomainError
	List() DomainError
}

type IController interface {
	HandleRequest(request string) (string, error)
	HandleGet(key string) (string, error)
	HandlePut(key string, value string) (string, error)
	HandleDelete(key string) (string, error)
	HandleList() (string, error)
}

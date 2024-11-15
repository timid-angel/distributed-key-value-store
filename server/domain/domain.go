package domain

type IDomainError interface {
	error
}

type DomainError struct {
	message string
}

func (err DomainError) Error() string {
	return err.message
}

func NewDomainError(message string) DomainError {
	return DomainError{message: message}
}

type IService interface {
	Get(key string) (string, IDomainError)
	Put(key string, value string) IDomainError
	Delete(key string) IDomainError
	List() (map[string]string, IDomainError)
}

type IController interface {
	HandleRequest(request string) string
	HandleGet(key string) string
	HandlePut(key string, value string) string
	HandleDelete(key string) string
	HandleList() string
}

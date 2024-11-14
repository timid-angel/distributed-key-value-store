package service

import "distributed-key-value-store/server/domain"

type Service struct{}

func (service *Service) Get(key string) domain.DomainError {
	panic("not implemented") // TODO: Implement
}

func (service *Service) Put(key string, value string) domain.DomainError {
	panic("not implemented") // TODO: Implement
}

func (service *Service) Delete(key string) domain.DomainError {
	panic("not implemented") // TODO: Implement
}

func (service *Service) List() domain.DomainError {
	panic("not implemented") // TODO: Implement
}

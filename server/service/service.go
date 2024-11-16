package service

import (
	"distributed-key-value-store/server/domain"
	"fmt"

	"github.com/gocql/gocql"
)

type Service struct {
	keyspace string
	session  *gocql.Session
}

func NewService(session *gocql.Session, keyspace string) *Service {
	return &Service{
		session:  session,
		keyspace: keyspace,
	}
}

func (service *Service) Get(key string) (string, domain.IDomainError) {
	var value string
	if err := service.session.Query(fmt.Sprintf(`SELECT value FROM %v.store WHERE key = ?`, service.keyspace), key).Scan(&value); err != nil {
		if err == gocql.ErrNotFound {
			return "", domain.NewDomainError(fmt.Sprintf("failed to GET key '%v': key not found", key))
		}

		return "", domain.NewDomainError(fmt.Sprintf("failed to GET key '%v': '%v'", key, err))
	}

	return value, nil
}

func (service *Service) Put(key string, value string) domain.IDomainError {
	err := service.session.Query(fmt.Sprintf("INSERT INTO %v.store (key, value) VALUES (?, ?)", service.keyspace), key, value).Exec()
	if err != nil {
		return domain.NewDomainError(fmt.Sprintf("failed to PUT '%v' to '%v': %v", key, value, err))
	}

	return nil
}

func (service *Service) Delete(key string) domain.IDomainError {
	if _, err := service.Get(key); err != nil {
		return domain.NewDomainError("failed to delete entry: key does not exist")
	}

	err := service.session.Query(fmt.Sprintf("DELETE FROM %v.store WHERE KEY = ?", service.keyspace), key).Exec()
	if err != nil {
		return domain.NewDomainError(fmt.Sprintf("failed to delete entry with key '%v': %v", key, err))
	}

	return nil
}

func (service *Service) List() (map[string]string, domain.IDomainError) {
	iter := service.session.Query(fmt.Sprintf(`SELECT key, value FROM %v.store`, service.keyspace)).Iter()
	results := make(map[string]string)
	var key, value string
	for iter.Scan(&key, &value) {
		results[key] = value
	}

	if err := iter.Close(); err != nil {
		return results, domain.NewDomainError(fmt.Sprintf("failed to fetch all data: %v", err))
	}

	return results, nil
}

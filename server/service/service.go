package service

import (
	"distributed-key-value-store/server/domain"
	"fmt"
	"log"
	"strings"

	"github.com/gocql/gocql"
)

type Service struct {
	keyspace string
	session  *gocql.Session
}

func (service *Service) Get(key string) (string, domain.IDomainError) {
	var value string
	if err := service.session.Query(fmt.Sprintf(`SELECT value FROM %v.store WHERE key = ?`, service.keyspace), key).Scan(&value); err != nil {
		if err == gocql.ErrNotFound {
			return "", domain.NewDomainError(fmt.Sprintf("failed to GET key '%v': key not found", key))
		}

		return "", domain.NewDomainError(fmt.Sprintf("failed to GET key '%v': %v", key, err))
	}

	return value, nil
}

func (service *Service) Put(key string, value string) domain.IDomainError {
	err := service.session.Query(fmt.Sprintf("INSERT INTO %v (key, value) VALUES (?, ?)", service.keyspace), key, value)
	if err != nil {
		return domain.NewDomainError(fmt.Sprintf("failed to PUT '%v' to '%v': %v", key, value, err))
	}

	return nil
}

func (service *Service) Delete(key string) domain.IDomainError {
	err := service.session.Query(fmt.Sprintf("DELETE FROM %v.store WHERE KEY = ?", service.keyspace), key)
	if err != nil {
		return domain.NewDomainError(fmt.Sprintf("failed to delete entry with key '%v': %v", key, err))
	}

	return nil
}

func (service *Service) List() string {
	iter := service.session.Query(fmt.Sprintf(`SELECT key, value FROM %v.store`, service.keyspace)).Iter()
	results := []string{}
	var key, value string
	for iter.Scan(&key, &value) {
		results = append(results, fmt.Sprintf("%v: %v", key, value))
	}

	if err := iter.Close(); err != nil {
		log.Fatalf("failed to fetch all data: %v", err)
	}

	return strings.Join(results, "; ")
}

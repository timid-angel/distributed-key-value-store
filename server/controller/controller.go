package controller

import (
	"distributed-key-value-store/server/domain"
	"fmt"
	"slices"
	"strings"
)

type Controller struct {
	service domain.IService
}

func NewController(service domain.IService) domain.IController {
	return &Controller{service: service}
}

func (controller *Controller) HandleRequest(request string) (string, error) {
	operations := []string{"get", "put", "delete", "list"}
	if request == "" {
		return "", fmt.Errorf("invalid command: only GET, PUT, DELETE and LIST operations are supported")
	}

	parts := strings.Split(request, " ")
	if len(parts) == 0 {
		return "", fmt.Errorf("invalid command: only GET, PUT, DELETE and LIST operations are supported")
	}

	operation := strings.ToLower(parts[0])
	if !slices.Contains(operations, operation) {
		return "", fmt.Errorf("invalid command: only GET, PUT, DELETE and LIST operations are supported")
	}

	if operation == "get" && len(parts) != 2 {
		return "", fmt.Errorf("invalid command: get requests must have the following syntax: `GET <KEY>`")
	}

	if operation == "delete" && len(parts) != 2 {
		return "", fmt.Errorf("invalid command: delete requests must have the following syntax: `DELETE <KEY>`")
	}

	if operation == "put" && len(parts) != 3 {
		return "", fmt.Errorf("invalid command: put requests must have the following syntax: `PUT <KEY> <VALUE>`")
	}

	if operation == "list" && len(parts) != 1 {
		return "", fmt.Errorf("invalid command: list requests must have the following syntax: `LIST`")
	}

	switch operation {
	case "get":
		return controller.HandleGet(parts[1])
	case "put":
		return controller.HandlePut(parts[1], parts[2])
	case "delete":
		return controller.HandleDelete(parts[1])
	case "list":
		return controller.HandleList()
	default:
	}

	return "", fmt.Errorf("invalid command: only GET, PUT, DELETE and LIST operations are supported")
}

func (controller *Controller) HandleGet(key string) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (controller *Controller) HandlePut(key string, value string) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (controller *Controller) HandleDelete(key string) (string, error) {
	panic("not implemented") // TODO: Implement
}

func (controller *Controller) HandleList() (string, error) {
	panic("not implemented") // TODO: Implement
}

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

func (controller *Controller) HandleRequest(request string) string {
	operations := []string{"get", "put", "delete", "list"}
	if request == "" {
		return "invalid command: only GET, PUT, DELETE and LIST operations are supported"
	}

	divParts := strings.Split(request, " ")
	parts := []string{}
	for _, v := range divParts {
		v = strings.TrimSpace(v)
		if v != "" && v != " " {
			parts = append(parts, v)
		}
	}

	if len(parts) == 0 {
		return "invalid command: only GET, PUT, DELETE and LIST operations are supported"
	}

	operation := strings.ToLower(parts[0])
	if !slices.Contains(operations, operation) {
		return "invalid command: only GET, PUT, DELETE and LIST operations are supported"
	}

	if operation == "get" && len(parts) != 2 {
		return "invalid command: get requests must have the following syntax: `GET <KEY>`"
	}

	if operation == "delete" && len(parts) != 2 {
		return "invalid command: delete requests must have the following syntax: `DELETE <KEY>`"
	}

	if operation == "put" && len(parts) != 3 {
		return "invalid command: put requests must have the following syntax: `PUT <KEY> <VALUE>`"
	}

	if operation == "list" && len(parts) != 1 {
		return "invalid command: list requests must have the following syntax: `LIST`"
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
		return "invalid command: only GET, PUT, DELETE and LIST operations are supported"
	}
}

func (controller *Controller) HandleGet(key string) string {
	value, err := controller.service.Get(key)
	if err != nil {
		return "Error: " + err.Error()
	}

	return value
}

func (controller *Controller) HandlePut(key string, value string) string {
	err := controller.service.Put(key, value)
	if err != nil {
		return "Error: " + err.Error()
	}

	return fmt.Sprintf("Successfully assigned key '%v' to value '%v", key, value)
}

func (controller *Controller) HandleDelete(key string) string {
	err := controller.service.Delete(key)
	if err != nil {
		return "Error: " + err.Error()
	}

	return fmt.Sprintf("Successfully removed entry with key '%v'", key)
}

func (controller *Controller) HandleList() string {
	res, err := controller.service.List()
	if err != nil {
		return "Error: " + err.Error()
	}

	result := ""
	for k, v := range res {
		result += fmt.Sprintf("%v: %v; ", k, v)
	}

	return fmt.Sprintf("List: %v", result)
}

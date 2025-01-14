package handler

import "github.com/osamikoyo/matesome/internal/service"

type UserHandler struct {
	Service service.UserValidator
}

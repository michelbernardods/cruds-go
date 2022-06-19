package handler

import (
	"github.com/dimiro1/health"
)

type Health struct {
}

func (c *Health) Handle() health.Handler {
	handler := health.NewHandler()
	return handler
}

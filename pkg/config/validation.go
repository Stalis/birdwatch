package config

import (
	"github.com/go-playground/validator/v10"
)

const (
	ErrInvalidPort = "invalid port value"
	ErrValidation  = "validation error"
)

func validate(config *Config) error {
	validate := validator.New()
	return validate.Struct(config)
}

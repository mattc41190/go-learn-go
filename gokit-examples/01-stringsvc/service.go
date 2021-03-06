package main

import (
	"errors"
	"strings"
)

// StringService provides operations on strings
type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type stringService struct{}

func (stringService) Uppercase(s string) (string, error) {
	if len(s) < 1 {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(s string) int {
	return len(s)
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty String")

// ServiceMiddleware type is a func which decorates the service being passed in
type ServiceMiddleware func(StringService) StringService

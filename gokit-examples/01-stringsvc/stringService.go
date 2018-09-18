package main

import (
	"context"
	"errors"
	"strings"
)

type StringService interface {
	Uppercase(context.Context, string) (string, error)
	Count(context.Context, string) int
}

type stringService struct{}

func (stringService) Uppercase(c context.Context, s string) (string, error) {
	if len(s) < 1 {
		return "", ErrEmpty
	}
	return strings.ToUpper(s), nil
}

func (stringService) Count(c context.Context, s string) int {
	return len(s)
}

// ErrEmpty is returned when input string is empty
var ErrEmpty = errors.New("Empty String")

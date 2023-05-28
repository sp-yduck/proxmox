package api

import (
	"github.com/pkg/errors"
)

var (
	ErrNotAuthorized = errors.New("not authorized to access endpoint")
	ErrTimeout       = errors.New("the operation has timed out")
	ErrNotFound      = errors.New("unable to find the item you are looking for")
)

func IsNotAuthorized(err error) bool {
	return err == ErrNotAuthorized
}

func IsTimeout(err error) bool {
	return err == ErrTimeout
}

func IsNotFound(err error) bool {
	return err == ErrNotFound
}

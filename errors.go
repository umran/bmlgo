package bmlgo

import (
	"errors"
)

var (
	// ErrorAuthenticationFailed ...
	ErrorAuthenticationFailed = errors.New("Authentication Failed")
	// ErrorNotAuthenticated ...
	ErrorNotAuthenticated = errors.New("Not Authenticated")
	// ErrorUnexpectedCode ...
	ErrorUnexpectedCode = errors.New("Unexpected Code")
	// ErrorUnexpectedStatusCode ...
	ErrorUnexpectedStatusCode = errors.New("Unexpected Status Code")
	// ErrorCursorUnreachable ...
	ErrorCursorUnreachable = errors.New("Cursor Unreachable")
)

func interpretStatusCode(code int) error {
	switch code {
	case 401:
		return ErrorNotAuthenticated
	default:
		return ErrorUnexpectedStatusCode
	}
}

func interpretCode(code int) error {
	switch code {
	case 2:
		return ErrorAuthenticationFailed
	default:
		return ErrorUnexpectedCode
	}
}

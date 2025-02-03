package error

import (
	"errors"
	"fmt"
)

// Base type for gRPC errors
type GRPCError struct {
	Msg string
}

func (e *GRPCError) Error() string {
	return e.Msg
}
func (e *GRPCError) Unwrap() error {
	return errors.New(e.Msg)
}

/**
 * Throw is a helper function to wrap errors with custom errors
 * @param custom error
 * @param err error
 * @return error
 */
func Throw(custom error, err error) error {
	return fmt.Errorf("%w-%v", custom, err)
}

// Custom Errors
var ConnectionError = &GRPCError{Msg: "Connection Error"}
var GetAllError = &GRPCError{Msg: "Get All Error"}
var GetVersionsError = &GRPCError{Msg: "Get Versions Error"}
var StreamEOF = &GRPCError{Msg: "Streaming cycles Error"}
var GetDetailsError = &GRPCError{Msg: "Get Details Error"}

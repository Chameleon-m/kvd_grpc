package handlers

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/protoadapt"
)

var (
	// ID должен быть > 0
	ErrBadRequeestFieldId = errdetails.BadRequest_FieldViolation{
		Field:       "Id",
		Description: "ID must be greater than zero",
	}
)

// Для возврата детальной ошибки переданных данных
func StatusInvalidArgumentDetails(details ...protoadapt.MessageV1) error {
	errorStatus := status.New(codes.InvalidArgument, "Invalid information received")
	ds, err := errorStatus.WithDetails(details...)
	if err != nil {
		return errorStatus.Err()
	}

	return ds.Err()
}

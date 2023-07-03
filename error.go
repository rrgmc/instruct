package instruct

import (
	"fmt"
	"reflect"

	"github.com/RangelReale/instruct/coerce"
)

var (
	ErrCoerce = coerce.ErrUnsupported
)

// An ValuesNotUsedError is returned when some values were not used.
type ValuesNotUsedError struct {
	Operation string
}

func (e ValuesNotUsedError) Error() string {
	return fmt.Sprintf("some values were not used on operation '%s'", e.Operation)
}

// An InvalidDecodeError describes an invalid argument passed to Decode.
// (The argument to Decode must be a non-nil pointer.)
type InvalidDecodeError struct {
	Type reflect.Type
}

func (e *InvalidDecodeError) Error() string {
	if e.Type == nil {
		return "error: Decode(nil)"
	}

	if e.Type.Kind() != reflect.Pointer {
		return "error: Decode(non-pointer " + e.Type.String() + ")"
	}
	return "error: Decode(nil " + e.Type.String() + ")"
}

// A RequiredError is returned when some values were not used.
type RequiredError struct {
	IsStructOption bool
	Operation      string
	FieldName      string
	TagName        string
}

func (e RequiredError) Error() string {
	f := "field"
	if e.IsStructOption {
		f = "struct option"
	}

	return fmt.Sprintf("%s '%s' (tag name '%s') with operation '%s' is required but was not set",
		f, e.FieldName, e.TagName, e.Operation)
}

// A OperationNotSupportedError is returned when an operation is not supported on the field.
type OperationNotSupportedError struct {
	Operation string
	FieldName string
}

func (e OperationNotSupportedError) Error() string {
	return fmt.Sprintf("operation '%s' not supported (no field type, maybe struct option?) for field '%s'",
		e.Operation, e.FieldName)
}

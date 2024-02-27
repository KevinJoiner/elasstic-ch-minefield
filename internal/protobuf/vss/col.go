// Package vss provides col names and types for the vss table.
package vss

import (
	"errors"
	"fmt"

	"github.com/ClickHouse/ch-go/proto"
)

var ErrInvalidType = errors.New("invalid type")

// MustGetColumn returns a function that returns a new column of the given type.
// It panics if the type is invalid.
func MustGetColumn(colType proto.ColumnType) func() proto.Column {
	// call mustGetColumn to panic on package init if the type is invalid
	_ = mustInfer(colType)
	return func() proto.Column { return mustInfer(colType) }
}

func mustInfer(colType proto.ColumnType) proto.Column {
	var auto proto.ColAuto
	if err := auto.Infer(colType); err != nil {
		panic(err)
	}
	return auto.Data
}

// NewCol returns a new column of the given field.
func NewCol(fieldName string) proto.Column {
	return newColsFunc[fieldName]()
}

// Append appends the value to the column.
func Append[T any](val T, col proto.Column) error {
	colOf, ok := col.(proto.ColumnOf[T])
	if !ok {
		var t T
		return fmt.Errorf("%w: provided col does not support type %T", ErrInvalidType, t)
	}
	colOf.Append(val)
	return nil
}

// MustAppend appends the value to the column.
// It panics if the type is invalid.
func MustAppend[T any](val T, col proto.Column) {
	var t T
	if boolVal, ok := any(t).(bool); ok {
		uval := uint8(0)
		if boolVal {
			uval = 1
		}
		if err := Append(uval, col); err != nil {
			panic(err)
		}
		return
	}
	if err := Append(val, col); err != nil {
		panic(err)
	}
}

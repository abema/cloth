package boolconv

import (
	"errors"
	"reflect"
	"strconv"
)

// Bool is a wrap of byte.
type Bool byte

const (
	// False means false on Bool.
	False Bool = iota
	// True means true on Bool.
	True
)

// NewBool converts bool into Bool.
func NewBool(b bool) Bool {
	if b {
		return True
	}
	return False
}

// BtoB converts bytes into Bool.
func BtoB(b []byte) Bool {
	return Bool(b[0])
}

// NewBoolByInterface converts interface into Bool.
func NewBoolByInterface(i interface{}) (Bool, error) {

	t := reflect.TypeOf(i)
	val := reflect.ValueOf(i)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	if t.Name() == "Bool" {
		if b, ok := (val.Interface()).(Bool); ok {
			return b, nil
		}
	}

	switch val.Kind() {

	case reflect.Bool:
		return NewBool((val.Interface()).(bool)), nil

	case reflect.Uint8:
		return Bool((val.Interface()).(byte)), nil

	case reflect.String:
		if b, err := strconv.ParseBool((val.Interface()).(string)); err == nil {
			return NewBool(b), nil
		}

	case reflect.Slice:
		v := val.Index(0)
		if b, ok := (v.Interface()).(byte); ok {
			return Bool(b), nil
		}
	}

	return False, errors.New("unsupported type")
}

// Tob returns bool.
func (b Bool) Tob() bool {
	return b == True
}

// Bytes returns []byte.
func (b Bool) Bytes() []byte {
	return []byte{byte(b)}
}

// String returns string.
func (b Bool) String() string {
	if b == True {
		return "true"
	}
	return "false"
}

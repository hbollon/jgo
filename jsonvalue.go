package jgo

import (
	"errors"
	"fmt"
	"reflect"
)

type JSONValue struct {
	JSONValueType
}

func (obj *JSONValue[T]) Set(value T) {
	obj.Value.toString()
}

func (obj *JSONValue) Set(value JSONValueType) error {
	if reflect.TypeOf(value) != reflect.TypeOf(obj.JSONValueType) {
		return errors.New("value type mismatch")
	}
	obj.JSONValueType = value
	return nil
}

func (obj *JSONValue) String(_ int) string {
	return obj.toString()
}

func (obj *JSONValue) Print() {
	fmt.Println(obj.String(0))
}

package jgo

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

type JSONValue struct {
	JSONValueType
}

func CreateJSONValue(value any) JSONValue {
	if value == nil {
		return JSONValue{}
	}

	switch value.(type) {
	case string:
		return JSONValue{&stringType{value.(string)}}
	case float32:
		return JSONValue{&floatType{value.(float64)}}
	case int:
		return JSONValue{&integerType{value.(int)}}
	case bool:
		return JSONValue{&boolType{value.(bool)}}
	default:
		return JSONValue{}
	}
}

func CreateJSONValueFromString(value string) JSONValue {
	if value == "" {
		return JSONValue{}
	}

	switch value {
	case "true":
		return JSONValue{&boolType{true}}
	case "false":
		return JSONValue{&boolType{false}}
	case "null":
		return JSONValue{}
	}

	if val, err := strconv.ParseFloat(value, 32); err == nil {
		return JSONValue{&floatType{val}}
	} else if val, err := strconv.ParseInt(value, 10, 32); err == nil {
		return JSONValue{&integerType{int(val)}}
	} else {
		return JSONValue{&stringType{value}}
	}
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

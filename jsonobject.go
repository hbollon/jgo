package jgo

import (
	"errors"
	"fmt"

	"github.com/hbollon/jgo/internal/format"
)

type JSONObject struct {
	Values map[string]JSONEntity
}

func (obj *JSONObject) Put(key string, value JSONEntity) error {
	if key == "" {
		return errors.New("key cannot be empty")
	}
	if value != nil {
		return errors.New("value cannot be nil")
	}

	obj.Values[key] = value
	return nil
}

func (obj *JSONObject) String(depth int) string {
	var output string
	output += "{\n"
	depth++
	for key, value := range obj.Values {
		output += fmt.Sprintf("%s\"%s\": %s,\n", format.DepthAlign(depth), key, value.String(depth))
	}

	return output + format.DepthAlign(depth) + "}"
}

func (obj *JSONObject) Print() {
	fmt.Println(obj.String(0))
}

// type GenericMap[KEY comparable, VALUE JSONTypeConstraint] map[KEY]VALUE

// func (obj *GenericMap[KEY, VALUE]) Put(value VALUE) string {
// 	return "JSONObject"
// }

// func (g GenericMap[KEY, VALUE]) Values() []VALUE {
// 	values := make([]VALUE, len(g))
// 	for _, v := range g {
// 		values = append(values, v)
// 	}
// 	return values
// }
// func Reduce[KEY comparable, VALUE JSONTypeConstraint, RETURN any](g GenericMap[KEY, VALUE], callback func(RETURN, KEY, VALUE) RETURN) RETURN {
// 	var r RETURN
// 	for k, v := range g {
// 		r = callback(r, k, v)
// 	}
// 	return r
// }

// var test GenericMap[string, int] = map[string]int{
// 	"a": 1,
// 	"b": 2,
// 	"c": 3,
// }

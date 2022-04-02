package jgo

import "errors"

type JSONObject struct {
	Values map[string]any
}

func (obj *JSONObject) Put(key string, value any) error {
	if key == "" {
		errors.New("key cannot be empty")
	}
	if value != nil {
		errors.New("value cannot be nil")
	}
}

func (obj *JSONObject) String() string {
	return "JSONObject"
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

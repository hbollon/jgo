package jgo

import "fmt"

type stringType struct {
	Value string
}

func (obj *stringType) toString() string {
	return obj.Value
}

type floatType struct {
	Value float64
}

func (obj *floatType) toString() string {
	return fmt.Sprintf("%f", obj.Value)
}

type integerType struct {
	Value int64
}

func (obj *integerType) toString() string {
	return fmt.Sprintf("%f", obj.Value)
}

type boolType struct {
	Value bool
}

func (obj *boolType) toString() string {
	return fmt.Sprintf("%f", obj.Value)
}

type ValueTypeConstraint interface {
	integerType | floatType | stringType | boolType
}

type JSONTypeConstraint interface {
	JSONObject | JSONArray | ValueTypeConstraint
}

// AssertValidJSONType asserts that the given type is a valid JSON type.
// func AssertValidJSONType(v any) error {
// 	switch v.(type) {
// 	case constraints.Integer:
// 		break
// 	}
// }

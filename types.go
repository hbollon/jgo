package jgo

import "fmt"

type JSONEntity interface {
	Print()
	String(int) string
}

type JSONValueType interface {
	toString() string
}

type stringType struct {
	string
}

func (obj *stringType) toString() string {
	return obj.string
}

type floatType struct {
	float64
}

func (obj *floatType) toString() string {
	return fmt.Sprintf("%f", obj.float64)
}

type integerType struct {
	int64
}

func (obj *integerType) toString() string {
	return fmt.Sprintf("%d", obj.int64)
}

type boolType struct {
	bool
}

func (obj *boolType) toString() string {
	return fmt.Sprintf("%v", obj.bool)
}

// type ValueTypeConstraint interface {
// 	integerType | floatType | stringType | boolType
// }

// type JSONTypeConstraint interface {
// 	JSONObject | JSONArray | ValueTypeConstraint
// }

// AssertValidJSONType asserts that the given type is a valid JSON type.
// func AssertValidJSONType(v any) error {
// 	switch v.(type) {
// 	case constraints.Integer:
// 		break
// 	}
// }

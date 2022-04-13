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
	int
}

func (obj *integerType) toString() string {
	return fmt.Sprintf("%d", obj.int)
}

type boolType struct {
	bool
}

func (obj *boolType) toString() string {
	return fmt.Sprintf("%v", obj.bool)
}

package jgo

type JSONValue[T ValueTypeConstraint] struct {
	Value T
}

func (obj *JSONValue[T]) Set(value T) {
	obj.Value.toString()
}

func (obj *JSONValue[T]) String() string {
	var test T
	switch test.(type) {

	}
}

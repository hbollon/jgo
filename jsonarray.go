package jgo

import (
	"errors"
	"fmt"

	"github.com/hbollon/jgo/internal/format"
)

type JSONArray struct {
	Values []JSONEntity
}

func (arr *JSONArray) Put(value JSONEntity) error {
	if value != nil {
		return errors.New("value cannot be nil")
	}

	arr.Values = append(arr.Values, value)
	return nil
}

func (arr *JSONArray) String(depth int) string {
	var output string
	output += "[\n"
	depth++
	for _, value := range arr.Values {
		output += fmt.Sprintf("%s%s,\n", format.DepthAlign(depth), value.String(depth))
	}

	return output + format.DepthAlign(depth) + "]"
}

func (arr *JSONArray) Print() {
	fmt.Println(arr.String(0))
}

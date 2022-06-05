package jgo

import (
	"errors"
	"fmt"
)

func Unmarshal(input any) (JSONObject, error) {
	switch input.(type) {
	case string:
		return UnmarshalFromString(input.(string))
	case JSONTokenizer:
		tokenizer := input.(JSONTokenizer)
		return UnmarshalFromTokenizer(&tokenizer)
	case map[string]any:
		return UnmarshalFromGenericMap(input.(map[string]any)), nil
	case map[string]JSONEntity:
		return UnmarshalFromMap(input.(map[string]JSONEntity)), nil
	default:
		return JSONObject{}, errors.New("Unmarshal: Unsupported type")
	}
}

func UnmarshalFromString(input string) (JSONObject, error) {
	if input == "" {
		return JSONObject{}, errors.New("UnmarshalFromString: Empty input")
	}
	tokenizer := NewJSONTokenizer(input)
	return UnmarshalFromTokenizer(tokenizer)
}

func UnmarshalFromTokenizer(t *JSONTokenizer) (JSONObject, error) {
	if t.Eof {
		return JSONObject{}, errors.New("UnmarshalFromTokenizer: EOF")
	}

	var c rune
	var key string
	obj := JSONObject{}

	// Check if the first character is a '{'
	if val, err := t.NextCharacter(); val != '{' || err != nil {
		if err != nil {
			return JSONObject{}, err
		}
		return JSONObject{}, errors.New("UnmarshalFromTokenizer: Expected '{'")
	}
	for {
		prevCharacter := t.Previous
		if val, err := t.NextCharacter(); err != nil {
			return JSONObject{}, err
		} else {
			c = val
		}

		switch c {
		case 0:
			return JSONObject{}, errors.New("UnmarshalFromTokenizer: EOF")
		case '}':
			return obj, nil
		case '{':
			if prevCharacter == '{' {
				return JSONObject{}, errors.New("UnmarshalFromTokenizer: Nested JSON objects inside another object are not allowed")
			}
		case '[':
			if prevCharacter == '{' {
				return JSONObject{}, errors.New("UnmarshalFromTokenizer: Nested JSON arrays inside an object are not allowed")
			}
		default:
			if err := t.Back(); err != nil {
				return JSONObject{}, err
			}
			if tmp, err := t.NextValue(); err != nil {
				return JSONObject{}, err
			} else {
				key = tmp.String(0)
			}
		}

		c, err := t.NextCharacter()
		if err != nil {
			return JSONObject{}, err
		}
		if c != ':' {
			return JSONObject{}, errors.New("UnmarshalFromTokenizer: Expected ':'")
		}

		if key != "" {
			if obj.Get(key) != nil {
				return JSONObject{}, errors.New("UnmarshalFromTokenizer: Duplicate key")
			}

			val, err := t.NextValue()
			if err != nil {
				return JSONObject{}, err
			}
			if val != nil {
				obj.Put(key, val)
			}
		}

		c, err = t.NextCharacter()
		if err != nil {
			return JSONObject{}, err
		}
		switch c {
		case ',':
			if c, err = t.NextCharacter(); err != nil {
				return JSONObject{}, err
			} else if c == '}' {
				return obj, nil
			}
			t.Back()
			break
		case '}':
			return obj, nil
		default:
			return JSONObject{}, errors.New(fmt.Sprintf("UnmarshalFromTokenizer: Unexpected character '%c' at line %d", c, t.Line))
		}
	}
}

func UnmarshalFromGenericMap(input map[string]any) JSONObject {
	return JSONObject{}
}

func UnmarshalFromMap(input map[string]JSONEntity) JSONObject {
	return JSONObject{}
}

func UnmarshalJSONArrayFromTokenizer(t *JSONTokenizer) (JSONArray, error) {
	if t.Eof {
		return JSONArray{}, errors.New("UnmarshalJSONArrayFromTokenizer: EOF")
	}

	var c rune
	var err error
	arr := JSONArray{}

	// Check if the first character is a '['
	if val, err := t.NextCharacter(); val != '[' || err != nil {
		if err != nil {
			return JSONArray{}, err
		}
		return JSONArray{}, errors.New("UnmarshalJSONArrayFromTokenizer: Expected '['")
	}

	c, err = t.NextCharacter()
	if err != nil {
		return JSONArray{}, err
	}

	if c == 0 {
		return JSONArray{}, errors.New("UnmarshalJSONArrayFromTokenizer: Array is unclosed")
	}
	if c == ']' {
		return JSONArray{}, nil
	}
	if c != ']' && t.Back() != nil {
		return JSONArray{}, errors.New("UnmarshalJSONArrayFromTokenizer: Back failed l.147")
	}
	for {
		prevCharacter := t.Previous
		val, err := t.NextCharacter()
		if err != nil {
			return JSONArray{}, err
		}

		if val == ',' {
			if t.Back() != nil {
				return JSONArray{}, errors.New("UnmarshalJSONArrayFromTokenizer: Back failed l.159")
			}
			arr.Put(&JSONObject{})
		} else {
			if t.Back() != nil {
				return JSONArray{}, errors.New("UnmarshalJSONArrayFromTokenizer: Back failed l.164")
			}
			entity, err := t.NextValue()
			if err != nil {
				return JSONArray{}, err
			}
			if err = arr.Put(entity); err != nil {
				return JSONArray{}, err
			}
		}

		val, err = t.NextCharacter()
		if err != nil {
			return JSONArray{}, err
		}

		switch val {
		case 0:
			return JSONArray{}, errors.New("UnmarshalJSONArrayFromTokenizer: EOF")
		case ']':

			return arr, nil
		case '{':
			if prevCharacter == '{' {
				return JSONArray{}, errors.New("UnmarshalJSONArrayFromTokenizer: Unexpected '{'")
			}
		case '[':
			if prevCharacter == '{' {
				return JSONArray{}, errors.New("UnmarshalJSONArrayFromTokenizer: Unexpected '['")
			}
		case ',':
			c, err = t.NextCharacter()
			if err != nil {
				return JSONArray{}, err
			}
			if c == 0 {
				return JSONArray{}, errors.New("UnmarshalJSONArrayFromTokenizer: EOF")
			}
			if c == ']' {
				return JSONArray{}, errors.New("UnmarshalJSONArrayFromTokenizer: Unexpected ']' after ','")
			}
			if t.Back() != nil {
				return JSONArray{}, errors.New("UnmarshalJSONArrayFromTokenizer: Back failed l.199")
			}
		default:
			return JSONArray{}, errors.New(fmt.Sprintf("UnmarshalJSONArrayFromTokenizer: Unexpected character '%c' at line %d", c, t.Line))
		}
	}
}

func UnmarshalJSONArrayFromString(input string) (JSONArray, error) {
	if input == "" {
		return JSONArray{}, errors.New("UnmarshalJSONArrayFromString: Empty input")
	}
	tokenizer := NewJSONTokenizer(input)
	return UnmarshalJSONArrayFromTokenizer(tokenizer)
}

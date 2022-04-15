package jgo

import (
	"errors"
)

func Unmarshall(input any) (JSONObject, error) {
	switch input.(type) {
	case string:
		return UnmarshallFromString(input.(string))
	case JSONTokenizer:
		tokenizer := input.(JSONTokenizer)
		return UnmarshallFromTokenizer(&tokenizer)
	case map[string]any:
		return UnmarshallFromGenericMap(input.(map[string]any)), nil
	case map[string]JSONEntity:
		return UnmarshallFromMap(input.(map[string]JSONEntity)), nil
	default:
		return JSONObject{}, errors.New("Unmarshall: Unsupported type")
	}
}

func UnmarshallFromString(input string) (JSONObject, error) {
	if input == "" {
		return JSONObject{}, errors.New("UnmarshallFromString: Empty input")
	}
	tokenizer := NewJSONTokenizer(input)
	return UnmarshallFromTokenizer(tokenizer)
}

func UnmarshallFromTokenizer(t *JSONTokenizer) (JSONObject, error) {
	if t.Eof {
		return JSONObject{}, errors.New("UnmarshallFromTokenizer: EOF")
	}

	var c rune
	var key string
	obj := JSONObject{}

	// Check if the first character is a '{'
	if val, err := t.NextCharacter(); val != '{' || err != nil {
		if err != nil {
			return JSONObject{}, err
		}
		return JSONObject{}, errors.New("UnmarshallFromTokenizer: Expected '{'")
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
			return JSONObject{}, errors.New("UnmarshallFromTokenizer: EOF")
		case '}':
			return obj, nil
		case '{':
			if prevCharacter == '{' {
				return JSONObject{}, errors.New("UnmarshallFromTokenizer: Unexpected '{'")
			}
		case '[':
			if prevCharacter == '{' {
				return JSONObject{}, errors.New("UnmarshallFromTokenizer: Unexpected '['")
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
			return JSONObject{}, errors.New("UnmarshallFromTokenizer: Expected ':'")
		}

		if key != "" {
			if obj.Get(key) != nil {
				return JSONObject{}, errors.New("UnmarshallFromTokenizer: Duplicate key")
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
			return JSONObject{}, errors.New("UnmarshallFromTokenizer: Unexpected character")
		}
	}
}

func UnmarshallFromGenericMap(input map[string]any) JSONObject {
	return JSONObject{}
}

func UnmarshallFromMap(input map[string]JSONEntity) JSONObject {
	return JSONObject{}
}

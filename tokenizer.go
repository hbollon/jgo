package jgo

import (
	"bufio"
	"errors"
	"strings"

	"github.com/hbollon/jgo/internal/format"
)

type JSONTokenizer struct {
	// CurrentCharacterPosition is the current read character position on the current line.
	CurrentCharacterPosition int64
	// Eof indicate if the end of the input has been found.
	Eof bool
	// Index indicate the current read index of the input.
	Index int64
	// Line indicate the current line of the input.
	Line int64
	// Previous is the previous character read from the input.
	Previous rune
	// Reader for the input.
	Reader *bufio.Reader
	// UsePrevious is a flag to indicate that a previous character was requested.
	UsePrevious bool
	// CharactersPreviousLine is the number of characters read in the previous line.
	CharactersPreviousLine int64
}

func NewJSONTokenizer(input string) *JSONTokenizer {
	return &JSONTokenizer{
		CurrentCharacterPosition: 0,
		Eof:                      false,
		Index:                    0,
		Line:                     1,
		Previous:                 0,
		Reader:                   bufio.NewReader(strings.NewReader(input)),
		UsePrevious:              false,
		CharactersPreviousLine:   0,
	}
}

// Next returns the next character from the input.
// If the end of the input is reached, EOF is returned.
func (t *JSONTokenizer) Next() (rune, error) {
	if t.UsePrevious {
		t.UsePrevious = false
		return t.Previous, nil
	}
	if t.Eof {
		return 0, errors.New("Next: EOF")
	}

	r, _, err := t.Reader.ReadRune()
	if err != nil {
		t.Eof = true
		return 0, nil
	}
	t.incrementIndexes(r)
	t.Previous = r
	return r, nil
}

// NextCharacter returns the next character from the input excluding whitespaces.
func (t *JSONTokenizer) NextCharacter() (rune, error) {
	for {
		c, err := t.Next()
		if err != nil {
			return 0, err
		}

		if c == 0 || c > ' ' {
			return c, nil
		}
	}
}

func (t *JSONTokenizer) NextString() (string, error) {
	var sb strings.Builder
	for {
		c, err := t.Next()
		if err != nil {
			return "", err
		}

		switch c {
		case '"':
			return sb.String(), nil
		case '\\':
			c, err := t.Next()
			if err != nil {
				return "", err
			}
			if c, err = format.HandleDoubleEscapedCharacter(c); err != nil {
				return "", err
			} else {
				sb.WriteRune(c)
			}
			break
		case '\n':
			return "", errors.New("NextString: Unexpected newline")
		default:
			sb.WriteRune(c)
			break
		}
	}
}

func (t *JSONTokenizer) NextValue() (JSONEntity, error) {
	c, err := t.NextCharacter()
	if err != nil {
		return nil, err
	}

	switch c {
	case '"':
		str, err := t.NextString()
		if err != nil {
			return nil, err
		}
		return &JSONValue{&stringType{str}}, nil
	case '{':
		if err = t.Back(); err != nil {
			return nil, err
		}

		obj, err := UnmarshallFromTokenizer(t)
		return &obj, err
	// TODO: Add support for arrays
	// case '[':
	// 	return t.NextArray()
	default:
		break
	}

	var sb strings.Builder
	for c != ' ' && !strings.ContainsRune(",:]}/\\\"[{;=#", c) {
		sb.WriteRune(c)
		c, err = t.Next()
		if err != nil {
			return nil, err
		}
	}
	if !t.Eof {
		if err = t.Back(); err != nil {
			return nil, err
		}
	}

	val := strings.TrimSpace(sb.String())
	if val == "" {
		return nil, errors.New("NextValue: empty value")
	}

	valStr := CreateJSONValueFromString(val)
	return &valStr, nil
}

// Back allow the tokenizer to go back one character.
// If the tokenizer is at the beginning of the input or
// if the tokenizr is already set to use the previous character, an error is returned.
func (t *JSONTokenizer) Back() error {
	if t.UsePrevious {
		return errors.New("Back: Cannot back more than one character")
	}
	if t.Index == 0 {
		return errors.New("Back: Index is 0")
	}
	t.decrementIndexes()
	t.UsePrevious = true
	return nil
}

// Peek returns the next character from the input without advancing the tokenizer.
func (t *JSONTokenizer) Peek() (rune, error) {
	if t.UsePrevious {
		return t.Previous, nil
	}
	if t.Eof {
		return 0, errors.New("Peek: EOF")
	}

	r, _, err := t.Reader.ReadRune()
	if err != nil {
		return 0, err
	}
	t.Reader.UnreadRune()
	return r, nil
}

// incrementIndexes increments the indexes of the tokenizer using the character read.
func (t *JSONTokenizer) incrementIndexes(c rune) {
	if c > 0 {
		t.Index++
		t.CurrentCharacterPosition++
		if c == '\n' {
			t.Line++
			t.CharactersPreviousLine = t.CurrentCharacterPosition
			t.CurrentCharacterPosition = 0
		}
	}
}

// decrementIndexes decrements the indexes of the tokenizer based on the previuous character read.
func (t *JSONTokenizer) decrementIndexes() {
	t.Index--
	t.Eof = false
	if t.Previous == '\n' {
		t.Line--
		t.CurrentCharacterPosition = t.CharactersPreviousLine
	} else if t.CurrentCharacterPosition != 0 {
		t.CurrentCharacterPosition--
	}
}

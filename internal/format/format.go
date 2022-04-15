package format

import "errors"

func DepthAlign(depth int) string {
	var output string
	for i := 0; i < depth; i++ {
		output += "\t"
	}
	return output
}

func HandleDoubleEscapedCharacter(c rune) (rune, error) {
	switch c {
	case 'b':
		return '\b', nil
	case 'f':
		return '\f', nil
	case 't':
		return '\t', nil
	case 'n':
		return '\n', nil
	case 'r':
		return '\r', nil
	case '"':
	case '\'':
	case '\\':
	case '/':
		return c, nil
	default:
		return 0, errors.New("HandleDoubleEscapedCharacter: Unsupported character")
	}

	return 0, nil
}

package structural_patterns

import "strings"

type StringManipulator func(string) string

func ToLower(m StringManipulator) StringManipulator {
	return func(s string) string {
		lower := strings.ToLower(s)
		return m(lower)
	}
}

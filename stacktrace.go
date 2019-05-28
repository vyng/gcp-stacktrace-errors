package stacktrace

import "strings"

func Stacktrace(stack []byte, skip int) []byte {
	parts := strings.Split(string(stack), "\n")
	if len(parts) < 4 {
		return stack
	}
	b := strings.Builder{}
	b.WriteString(parts[0])
	parts = parts[3 + (skip * 2):]
	if len(parts) <= 0 {
		return stack
	}
	for _, s := range parts {
		b.WriteString("\n")
		b.WriteString(s)
	}

	return []byte(b.String())
}
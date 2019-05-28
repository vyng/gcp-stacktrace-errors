package stacktrace

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStacktrace(t *testing.T) {
	t.Run("stacktrace will return a proper stacktrace", testValidStacktrace())
	t.Run("stacktrace will return input if not at least 6 lines", testInvalidStacktrace())
}

func testValidStacktrace() func(t *testing.T) {
	return func(t *testing.T) {
		stack := `goroutine 1 [running]:
runtime/debug.Stack(0x1343e0, 0x40c108, 0x41a7a8, 0x1)
	/usr/local/go/src/runtime/debug/stack.go:24 +0xc0
main.main()
	/tmp/sandbox224636382/main.go:10 +0xa0`

		result := string(Stacktrace([]byte(stack), 0))
		expected := `goroutine 1 [running]:
main.main()
	/tmp/sandbox224636382/main.go:10 +0xa0`

		assert.Equalf(t, result == expected, true, "Expected properly created stacktrace.", result)
	}
}

func testInvalidStacktrace() func(t *testing.T) {
	return func(t *testing.T) {
		bytes := []byte{1, 2, 3}
		result := Stacktrace(bytes, 0)
		assert.Equal(t, string(bytes) == string(result), true, "Expected input to equal output")
	}
}

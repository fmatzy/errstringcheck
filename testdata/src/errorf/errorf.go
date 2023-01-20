package errorf

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.New("foo")
	str := "foo"

	fmt.Errorf("this is message")          // OK
	fmt.Errorf("this is message %q", str)  // OK
	fmt.Errorf("err=%s", err)              // want `invalid format for fmt.Errorf. Use "...: %v" or "...: %w" to format errors`
	fmt.Errorf("err: %s: suffix", err)     // want `invalid format for fmt.Errorf. Use "...: %v" or "...: %w" to format errors`
	fmt.Errorf("err: %s", err)             // want `invalid format for fmt.Errorf. Use "...: %v" or "...: %w" to format errors`
	fmt.Errorf("err=%v", err)              // want `invalid format for fmt.Errorf. Use "...: %v" or "...: %w" to format errors`
	fmt.Errorf("err: %v: suffix", err)     // want `invalid format for fmt.Errorf. Use "...: %v" or "...: %w" to format errors`
	fmt.Errorf("err: %v", err)             // OK
	fmt.Errorf("err=%w", err)              // want `invalid format for fmt.Errorf. Use "...: %v" or "...: %w" to format errors`
	fmt.Errorf("err: %w: suffix", err)     // want `invalid format for fmt.Errorf. Use "...: %v" or "...: %w" to format errors`
	fmt.Errorf("err: %w", err)             // OK
	fmt.Errorf("err=%s, str=%s", err, str) // want `invalid format for fmt.Errorf. Use "...: %v" or "...: %w" to format errors`

	// error.Error()
	fmt.Errorf("err=%s", err.Error())  // want `invalid format for fmt.Errorf. Use "...: %v" or "...: %w" to format errors`
	fmt.Errorf("err=%v", err.Error())  // want `invalid format for fmt.Errorf. Use "...: %v" or "...: %w" to format errors`
	fmt.Errorf("err: %s", err.Error()) // want `invalid format for fmt.Errorf. Use "...: %v" or "...: %w" to format errors`
	fmt.Errorf("err: %v", err.Error()) // OK
}

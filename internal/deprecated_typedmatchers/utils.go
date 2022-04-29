package typedmatchers

import "fmt"

func formatGottenArg[T any](m Matcher[T], arg interface{}) string {
	switch actual := arg.(type) {
	case T:
		if gs, ok := m.(GotFormatter[T]); ok {
			return gs.Got(actual)
		}
	case *T:
		if gs, ok := m.(GotFormatter[T]); ok {
			return gs.Got(*actual)
		}
	}
	return fmt.Sprintf("%+v (%T)", arg, arg)
}

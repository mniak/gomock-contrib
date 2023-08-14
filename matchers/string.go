package matchers

import (
	"fmt"
	"strings"
)

func HasPrefix(prefix string) inlineMatcher {
	return *Inline(fmt.Sprintf(`a string with prefix %q`, prefix), func(value any) bool {
		valueString, isString := value.(string)
		return isString && strings.HasPrefix(valueString, prefix)
	})
}

func HasSuffix(suffix string) inlineMatcher {
	return *Inline(fmt.Sprintf(`a string with suffix %q`, suffix), func(value any) bool {
		valueString, isString := value.(string)
		return isString && strings.HasSuffix(valueString, suffix)
	})
}

func EqualFold(other string) inlineMatcher {
	return *Inline(fmt.Sprintf(`equal to %q (case insensitive)`, other), func(value any) bool {
		valueString, isString := value.(string)
		return isString && strings.EqualFold(valueString, other)
	})
}

func Contains(substr string) inlineMatcher {
	return *Inline(fmt.Sprintf(`containing %q`, substr), func(value any) bool {
		valueString, isString := value.(string)
		return isString && strings.Contains(valueString, substr)
	})
}

func ContainsAny(prefix string) inlineMatcher {
	return *Inline(fmt.Sprintf(`containing any of the chars %q`, prefix), func(value any) bool {
		valueString, isString := value.(string)
		return isString && strings.ContainsAny(valueString, prefix)
	})
}

func ContainsRune(r rune) inlineMatcher {
	return *Inline(fmt.Sprintf(`contining the rune %q`, string(r)), func(value any) bool {
		valueString, isString := value.(string)
		return isString && strings.ContainsRune(valueString, r)
	})
}

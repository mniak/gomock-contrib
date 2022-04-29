package typedmatchers

import "github.com/golang/mock/gomock"

var (
	_ Matcher[any] = gomock.Any()
	_ Matcher[any] = gomock.Eq(nil)
)

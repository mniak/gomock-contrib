package typedmatchers

import "github.com/golang/mock/gomock"

var (
	_ Matcher[interface{}] = gomock.Any()
	_ Matcher[interface{}] = gomock.Eq(nil)
)

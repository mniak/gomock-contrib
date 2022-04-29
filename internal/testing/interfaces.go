package testing

import "github.com/golang/mock/gomock"

type GoMockMatcherAndGotFormatter interface {
	gomock.Matcher
	gomock.GotFormatter
}

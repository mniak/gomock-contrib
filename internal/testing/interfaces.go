package testing

import "github.com/golang/mock/gomock"

type MatcherGotFormatter interface {
	gomock.Matcher
	gomock.GotFormatter
}

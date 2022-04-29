package utils

import "github.com/golang/mock/gomock"

func ArgAsMatcher(arg any) gomock.Matcher {
	var submatcher gomock.Matcher
	if sub, is := arg.(gomock.Matcher); is {
		submatcher = sub
	} else {
		submatcher = gomock.Eq(arg)
	}
	return submatcher
}

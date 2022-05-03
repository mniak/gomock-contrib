package matchers

import (
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
)

func IsGUID() gomock.Matcher {
	return Inline("is valid GUID", func(arg any) bool {
		cid, isString := arg.(string)
		if !isString {
			return false
		}
		_, err := uuid.Parse(cid)
		return err == nil
	})
}

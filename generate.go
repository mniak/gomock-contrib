package gomockContrib

// Generate OpenTelemetry Mocks
//go:generate mockgen -package=mocks -destination=internal/testing/mocks/gomock_matcher.go "github.com/golang/mock/gomock" Matcher
//go:generate mockgen -package=mocks -destination=internal/testing/mocks/gomock_matcher_and_gotformatter.go "github.com/mniak/gomock-contrib/internal/testing" GoMockMatcherAndGotFormatter

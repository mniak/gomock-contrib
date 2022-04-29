package gomockContrib

// Generate OpenTelemetry Mocks
//go:generate mockgen -package=mocks -destination=internal/testing/mocks/gomock_matcher.go "github.com/golang/mock/gomock" Matcher

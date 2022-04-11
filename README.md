GoMock Contrib
================

Matchers for GoMock.
For now, the main purpose is to provide typed (generic) matchers.


Usage
-------


```go
import "github.com/mniak/gomock-contrib/matchers"

[...]

myMock.EXPECT().
    Method(matchers.Typed[string](
        typedmatchers.Inline("custom validation", func (value string) bool {
            //
        }
    )

```
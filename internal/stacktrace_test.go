package internal

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseAnyStackTrace(t *testing.T) {
	stack := `
goroutine 1 [running]:
runtime/debug.Stack()
        /home/stephan/sdk/go1.20.1/src/runtime/debug/stack.go:24 +0x65
erro.parseStackTrace(0x751710?)
        /home/stephan/Development/erro/regexp.go:35 +0x1e
erro.(*logger).Errorf(0x751710, {0x5df7f6, 0x19}, {0x62f120?, 0xc000016270}, {0xc000016280, 0x1, 0x1})
        /home/stephan/Development/erro/logger.go:136 +0xb0
erro.Errorf(...)
        /home/stephan/Development/erro/erro.go:71
main.someBigFunction(0x6303f8?)
        /home/stephan/Development/erro/examples/error.go:37 +0x125
main.wrapingFunc(...)
        /home/stephan/Development/erro/examples/error.go:24
main.main()
        /home/stephan/Development/erro/examples/error.go:19 +0x47d
`
	items := parseAnyStackTrace(stack, 1)
	//for _, i := range items {
	//	spew.Dump(i)
	//}
	assert.Equal(t, 4, len(items))
}

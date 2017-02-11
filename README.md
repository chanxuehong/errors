# Forked from https://github.com/pkg/errors

**go1.7+ required**

## Usage

```go
package main

import (
	stderrors "errors"
	"log"

	"github.com/chanxuehong/errors"
)

func main() {
	err := test3()
	if err != nil {
		log.Println(errors.String(err))
		return
	}
}

func test0() error {
	return stderrors.New("original message")
}

func test1() error {
	err := test0()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func test2() error {
	err := test1()
	if err != nil {
		return errors.Wrap(test1(), "test2 wrap message")
	}
	return nil
}

func test3() error {
	err := test2()
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
```

result:
```
2017/02/11 13:48:49 test2 wrap message: original message
main.test1
        /Users/chan/gopath/src/test1/main.go:25
main.test2
        /Users/chan/gopath/src/test1/main.go:33
main.test3
        /Users/chan/gopath/src/test1/main.go:39
main.main
        /Users/chan/gopath/src/test1/main.go:11
runtime.main
        /usr/local/go/src/runtime/proc.go:185
runtime.goexit
        /usr/local/go/src/runtime/asm_amd64.s:2197
```

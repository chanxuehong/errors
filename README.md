# Forked from https://github.com/pkg/errors

**go1.7+ required**

## Usage

```go
package main

import (
	stderrors "errors"
	"fmt"

	"github.com/chanxuehong/errors"
)

func main() {
	err := test3()
	if err != nil {
		fmt.Println(errors.ErrorStack(err))
		return
	}
}

func test0() error {
	return stderrors.New("original message")
}

func test1() error {
	err := test0()
	if err != nil {
		return errors.Wrap(err)
	}
	return nil
}

func test2() error {
	err := test1()
	if err != nil {
		return errors.Wrap(err, "test2 wrap message")
	}
	return nil
}

func test3() error {
	err := test2()
	if err != nil {
		return errors.Wrap(err)
	}
	return nil
}
```

The result is:
```
original message
main.test1
	test2/main.go:25
main.test2
	test2/main.go:31
main.test3
	test2/main.go:39
main.main
	test2/main.go:11
test2 wrap message
```

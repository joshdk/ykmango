[![License](https://img.shields.io/github/license/joshdk/ykmango.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/joshdk/ykmango?status.svg)](https://godoc.org/github.com/joshdk/ykmango)
[![Go Report Card](https://goreportcard.com/badge/github.com/joshdk/ykmango)](https://goreportcard.com/report/github.com/joshdk/ykmango)
[![CircleCI](https://circleci.com/gh/joshdk/ykmango.svg?&style=shield)](https://circleci.com/gh/joshdk/ykmango/tree/master)

# YKManGo

🔑 Prompt a YubiKey device to generate an OATH code

## Installing

You can fetch this library by running the following

    go get -u github.com/joshdk/ykmango

## Usage

```go
import (
	"fmt"
	"github.com/joshdk/ykmango"
)

// List the currently configured OATH slot names.
names, err := ykman.List()
if err != nil {
	panic(err.Error())
}

for _, name := range names {
	fmt.Printf("Found code named: %s\n", name)
	// Found code named: aws-mfa
}
```

## License

This library is distributed under the [MIT License](https://opensource.org/licenses/MIT), see LICENSE.txt for more information.
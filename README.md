[![License](https://img.shields.io/github/license/joshdk/ykmango.svg)](https://opensource.org/licenses/MIT)
[![GoDoc](https://godoc.org/github.com/joshdk/ykmango?status.svg)](https://godoc.org/github.com/joshdk/ykmango)
[![Go Report Card](https://goreportcard.com/badge/github.com/joshdk/ykmango)](https://goreportcard.com/report/github.com/joshdk/ykmango)
[![CircleCI](https://circleci.com/gh/joshdk/ykmango.svg?&style=shield)](https://circleci.com/gh/joshdk/ykmango/tree/master)
[![CodeCov](https://codecov.io/gh/joshdk/ykmango/branch/master/graph/badge.svg)](https://codecov.io/gh/joshdk/ykmango)

# YKManGo

🔑 Prompt a [YubiKey device](https://en.wikipedia.org/wiki/YubiKey) to generate an OATH code

## Installing

You can fetch this library by running the following

    go get -u github.com/joshdk/ykmango
    
Additionally, this library has a runtime dependency on [`ykman`](https://github.com/Yubico/yubikey-manager/tree/master/ykman) which must [be installed](https://developers.yubico.com/yubikey-manager/) before use.

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

// Generate an OATH code using the given slot name.
// You may need to touch your YubiKey device if the
// slot is configured to require touch.
code, err := ykman.Generate("aws-mfa")
if err != nil {
	panic(err.Error())
}

fmt.Printf("Your code is: %s\n", code)
// Your code is: 150509
```

## License

This library is distributed under the [MIT License](https://opensource.org/licenses/MIT), see LICENSE.txt for more information.
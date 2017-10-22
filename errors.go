// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package ykman

import (
	"errors"
)

var (
	ErrorYkmanNotFound = errors.New("ykman executable not found in $PATH")
	ErrorNotDetected   = errors.New("no yubikey detected")
	ErrorAborted       = errors.New("aborted")
	ErrorRemoved       = errors.New("yubikey removed")
	ErrorTimeout       = errors.New("timed out")
	ErrorUnknownName   = errors.New("unknown slot title")
)

// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package ykman

import (
	"strings"
)

func process(body string) []string {

	lines := []string{}

	for _, line := range strings.Split(body, "\n") {
		line := strings.TrimSpace(line)
		if len(line) == 0 {
			continue
		}

		lines = append(lines, line)
	}

	return lines
}

func linesContain(lines []string, phrase string) bool {
	for _, line := range lines {
		if strings.Contains(line, phrase) {
			return true
		}
	}

	return false
}

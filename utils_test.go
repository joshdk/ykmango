// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package ykman

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProcess(t *testing.T) {

	tests := []struct {
		name     string
		body     string
		expected []string
	}{
		{
			name:     "empty body",
			expected: []string{},
		},
		{
			name:     "single space",
			body:     " ",
			expected: []string{},
		},
		{
			name:     "multiple spaces",
			body:     "   ",
			expected: []string{},
		},
		{
			name:     "single tab",
			body:     "\t",
			expected: []string{},
		},
		{
			name:     "multiple tabs",
			body:     "\t\t\t",
			expected: []string{},
		},
		{
			name:     "mixed whitespace",
			body:     "\t \t \t ",
			expected: []string{},
		},
		{
			name: "single word",
			body: "alice",
			expected: []string{
				"alice",
			},
		},
		{
			name: "single word surrounded with whitespace",
			body: "\t \t \t alice\t \t \t ",
			expected: []string{
				"alice",
			},
		},
		{
			name: "multiple words",
			body: "alice bob carol",
			expected: []string{
				"alice bob carol",
			},
		},
		{
			name: "multiple words surrounded with whitespace",
			body: "\t \t \t alice bob carol\t \t \t ",
			expected: []string{
				"alice bob carol",
			},
		},
		{
			name: "single line ending with a newline",
			body: "alice bob carol\n",
			expected: []string{
				"alice bob carol",
			},
		},
		{
			name: "multiple body",
			body: `
				alice bob carol
				dave eve fred
			    grant henry ida
			`,
			expected: []string{
				"alice bob carol",
				"dave eve fred",
				"grant henry ida",
			},
		},
		{
			name: "multiple body some blank",
			body: `
				alice bob carol

				dave eve fred

			    grant henry ida
			`,
			expected: []string{
				"alice bob carol",
				"dave eve fred",
				"grant henry ida",
			},
		},
		{
			name: "multiple blank body",
			body: `



			`,
			expected: []string{},
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.name)

		t.Run(name, func(t *testing.T) {
			actual := process(test.body)

			assert.Equal(t, test.expected, actual)
		})
	}

}

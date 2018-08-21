// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package ykman

import (
	"errors"
	"fmt"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListParse(t *testing.T) {

	var errorGeneric = errors.New("generic error")

	tests := []struct {
		title         string
		body          string
		error         error
		expectedNames []string
		expectedError error
	}{
		{
			title:         "empty results",
			expectedNames: []string{},
		},
		{
			title: "single result",
			body: `
				aws
			`,
			expectedNames: []string{"aws"},
		},
		{
			title: "single result with spaces",
			body: `
				aws main account
			`,
			expectedNames: []string{"aws main account"},
		},
		{
			title: "multiple results",
			body: `
				aws
				aws-cn
				aws-us-gov
			`,
			expectedNames: []string{"aws", "aws-cn", "aws-us-gov"},
		},
		{
			title: "multiple results with spaces",
			body: `
				aws
				aws-cn
				aws main account
			`,
			expectedNames: []string{"aws", "aws-cn", "aws main account"},
		},
		{
			title: "ykman executable missing",
			error: &exec.Error{
				Err: exec.ErrNotFound,
			},
			expectedError: ErrorYkmanNotFound,
		},
		{
			title: "no yubikey detected",
			body: `
				Usage: ykman [OPTIONS] COMMAND [ARGS]...
				Error: No YubiKey detected!
			`,
			error:         errorGeneric,
			expectedError: ErrorYubikeyNotDetected,
		},

		{
			title: "no yubikey detected",
			body: `
				Usage: ykman [OPTIONS] COMMAND [ARGS]...
				Error: Failed connecting to the YubiKey.
			`,
			error:         errorGeneric,
			expectedError: ErrorYubikeyNotDetected,
		},
		{
			title:         "generic error",
			error:         errorGeneric,
			expectedError: errorGeneric,
		},
	}

	for index, test := range tests {

		name := fmt.Sprintf("case #%d - %s", index, test.title)

		t.Run(name, func(t *testing.T) {
			actualNames, actualError := parseList(test.body, test.error)

			assert.Equal(t, test.expectedError, actualError)

			assert.Equal(t, test.expectedNames, actualNames)
		})
	}

}

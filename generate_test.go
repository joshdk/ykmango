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

func TestGenerateParse(t *testing.T) {

	var errorGeneric = errors.New("generic error")

	tests := []struct {
		title         string
		name          string
		body          string
		error         error
		expectedError error
		expectedCode  string
	}{
		{
			title:         "empty results",
			expectedError: ErrorSlotNameUnknown,
		},
		{
			title: "single result",
			name:  "aws",
			body: `
				aws  123456
			`,
			expectedCode: "123456",
		},
		{
			title: "single long result",
			name:  "aws",
			body: `
				aws  12345678
			`,
			expectedCode: "12345678",
		},
		{
			title: "multiple results",
			name:  "aws",
			body: `
				aws  123456
				aws  867530
			`,
			expectedCode: "123456",
		},
		{
			title: "code too short",
			name:  "aws",
			body: `
				aws  1234
			`,
			expectedError: ErrorSlotNameUnknown,
		},
		{
			title: "bad code format",
			name:  "aws",
			body: `
				aws  DEADBEEF
			`,
			expectedError: ErrorSlotNameUnknown,
		},
		{
			title: "different name returned",
			name:  "aws",
			body: `
				aws-us-gov  123456
			`,
			expectedError: ErrorSlotNameUnknown,
		},
		{
			title: "error ykman executable missing",
			error: &exec.Error{
				Err: exec.ErrNotFound,
			},
			expectedError: ErrorYkmanNotFound,
		},
		{
			title: "error no yubikey detected",
			body: `
				Usage: ykman [OPTIONS] COMMAND [ARGS]...

				Error: Failed connecting to the YubiKey.
			`,
			error:         errorGeneric,
			expectedError: ErrorYubikeyNotDetected,
		},
		{
			title: "error aborted",
			body: `
				Touch your YubiKey...
				^C
				Aborted!
			`,
			error:         errorGeneric,
			expectedError: ErrorYkmanInterrupted,
		},
		{
			title: "error yubikey removed",
			body: `
				Touch your YubiKey...
				Traceback (most recent call last):
				  File "/usr/local/bin/ykman", line 11, in <module>
					load_entry_point('yubikey-manager==0.4.5', 'console_scripts', 'ykman')()
				  <...snip...>
				  File "/usr/local/Cellar/ykman/0.4.5/libexec/lib/python2.7/site-packages/ykman/driver_ccid.py", line 123, in send_apdu
					raise CCIDError(e)
				ykman.driver_ccid.CCIDError: Failed to transmit with protocol T1. Transaction failed.
			`,
			error:         errorGeneric,
			expectedError: ErrorYubikeyRemoved,
		},
		{
			title: "error yubikey timed out",
			body: `
				Touch your YubiKey...
				Traceback (most recent call last):
				  File "/usr/local/bin/ykman", line 11, in <module>
					load_entry_point('yubikey-manager==0.4.5', 'console_scripts', 'ykman')()
				  <...snip...>
				  File "/usr/local/Cellar/ykman/0.4.5/libexec/lib/python2.7/site-packages/ykman/oath.py", line 203, in send_apdu
					raise APDUError(resp, sw)
				ykman.driver_ccid.APDUError: APDU error: SW=0x6982
			`,
			error:         errorGeneric,
			expectedError: ErrorYubikeyTimeout,
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
			actualCode, actualError := parseGenerate(test.body, test.error, test.name)

			assert.Equal(t, test.expectedError, actualError)

			assert.Equal(t, test.expectedCode, actualCode)
		})
	}

}

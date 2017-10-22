// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package ykman

import (
	"os/exec"
	"regexp"
)

func Generate(name string) (string, error) {

	cmd := exec.Command("ykman", "oath", "code", name)

	output, err := cmd.CombinedOutput()

	return parseGenerate(string(output), err, name)
}

func parseGenerate(body string, err error, name string) (string, error) {

	lines := process(body)

	if err != nil {
		// Check if this is an exec.Error stating that the ykman executable was not found
		if execErr, ok := err.(*exec.Error); ok {
			if execErr.Err == exec.ErrNotFound {
				return "", ErrorYkmanNotFound
			}
		}

		// Case where a YubiKey isn't plugged in
		if linesContain(lines, "Failed connecting to the YubiKey") {
			return "", ErrorNotDetected
		}

		// Case where ykman was killed/interruped with a signal
		if linesContain(lines, "Aborted!") {
			return "", ErrorAborted
		}

		// Case where ykman dumps a Python exception
		if linesContain(lines, "Traceback (most recent call last)") {
			// Case where yubikey is removed mid-operation
			if linesContain(lines, "Failed to transmit with protocol") {
				return "", ErrorRemoved
			}

			// Case where YubiKey was not touched in time
			if linesContain(lines, "APDU error") {
				return "", ErrorTimeout
			}
		}

		// Generic catch-all
		return "", err
	}

	oathCodeRegex := regexp.MustCompile("^" + name + "\\s+(\\d{6,})$")

	for _, line := range lines {
		matches := oathCodeRegex.FindStringSubmatch(line)

		if len(matches) == 2 {
			return matches[1], nil
		}
	}

	return "", ErrorUnknownName
}

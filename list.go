// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package ykman

import (
	"os/exec"
	"strings"
)

// List returns all currently configured OATH slot names
func List() ([]string, error) {

	cmd := exec.Command("ykman", "oath", "list")

	output, err := cmd.CombinedOutput()

	return parseList(string(output), err)
}

func parseList(body string, err error) ([]string, error) {

	lines := process(body)

	if err != nil {
		// Check if this is an exec.Error stating that the ykman executable was not found
		if execErr, ok := err.(*exec.Error); ok {
			if execErr.Err == exec.ErrNotFound {
				return nil, ErrorYkmanNotFound
			}
		}

		// Check the case where a YubiKey isn't plugged in
		if linesContain(lines, "No YubiKey detected") {
			return nil, ErrorYubikeyNotDetected
		}

		// Case where a YubiKey isn't plugged in
		if linesContain(lines, "Failed connecting to the YubiKey") {
			return nil, ErrorYubikeyNotDetected
		}

		// Generic catch-all
		return nil, err
	}

	names := make([]string, len(lines))

	for index, line := range lines {
		chunks := strings.Split(line, "\n")
		names[index] = chunks[0]
	}

	return names, nil
}

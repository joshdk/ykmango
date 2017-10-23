// Copyright 2017 Josh Komoroske. All rights reserved.
// Use of this source code is governed by an MIT-style
// license that can be found in the LICENSE.txt file.

package ykman

import (
	"errors"
)

var (
	// ErrorYkmanNotFound indicates that the ykman executable could not be found
	// Installation or $PATH configuration may be needed to correct
	ErrorYkmanNotFound = errors.New("ykman executable not found")

	// ErrorYkmanInterrupted indicates that the ykman process was killed with a signal
	ErrorYkmanInterrupted = errors.New("ykman interrupted")

	// ErrorYubikeyNotDetected indicates that a YubiKey is not currently plugged in
	ErrorYubikeyNotDetected = errors.New("yubikey not detected")

	// ErrorYubikeyRemoved indicates that a YubiKey was removed while in-use
	ErrorYubikeyRemoved = errors.New("yubikey removed")

	// ErrorYubikeyTimeout indicates that a YubiKey was not touched in time to generate an OATH code
	ErrorYubikeyTimeout = errors.New("yubikey timed out")

	// ErrorSlotNameUnknown indicates that the specified OATH slot name does not exist
	ErrorSlotNameUnknown = errors.New("slot name unknown")
)

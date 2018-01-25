// Copyright 2014 ISRG.  All rights reserved
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// Copied from https://github.com/letsencrypt/boulder/blob/master/test/test-tools.go
//
// See Q5 and Q11 here: https://www.mozilla.org/en-US/MPL/2.0/FAQ/ I think if
// we want to modify this file we have to release it publicly, otherwise we're
// fine.

package types

import (
	"testing"
)

// Assert a boolean
func assert(t *testing.T, result bool, message string) {
	t.Helper()
	if !result {
		t.Fatalf(message)
	}
}

// AssertNotError checks that err is nil
func assertNotError(t *testing.T, err error, message string) {
	t.Helper()
	if err != nil {
		t.Fatalf("%s: %s", message, err)
	}
}

// AssertError checks that err is non-nil
func assertError(t *testing.T, err error, message string) {
	t.Helper()
	if err == nil {
		t.Fatalf("%s: expected error but received none", message)
	}
}

// AssertEquals uses the equality operator (==) to measure one and two
func assertEquals(t *testing.T, one interface{}, two interface{}) {
	t.Helper()
	if one != two {
		t.Fatalf("[%v] != [%v]", one, two)
	}
}

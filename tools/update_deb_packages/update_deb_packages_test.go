package main

import (
	"testing"
	"testing/quick"
)

func TestAppendUniqLength(t *testing.T) {
	// appending a string either increases the length by one or it stays the same
	lengthCheck := func(slice []string, v string) bool {
		return len(appendUniq(slice, v)) == len(slice) || len(appendUniq(slice, v)) == len(slice)+1
	}

	if err := quick.Check(lengthCheck, nil); err != nil {
		t.Error(err)
	}
}

func TestAppendUniqMember(t *testing.T) {
	// after running appendUniq, the value can be found in the list
	memberCheck := func(slice []string, v string) bool {
		slice = appendUniq(slice, v)
		for _, x := range slice {
			if x == v {
				return true
			}
		}
		return false
	}

	if err := quick.Check(memberCheck, nil); err != nil {
		t.Error(err)
	}
}

func TestAppendUniqUniqueness(t *testing.T) {
	// the appended value is only once in the list
	uniquenessCheck := func(input []string) bool {
		m := make(map[string]bool)
		var s []string

		for _, x := range input {
			m[x] = true
			s = appendUniq(s, x)
		}
		if len(m) != len(s) {
			return false
		}
		return true
	}

	if err := quick.Check(uniquenessCheck, nil); err != nil {
		t.Error(err)
	}
}

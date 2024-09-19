package main

import "testing"

func TestIsSubsequence(t *testing.T) {
	var res bool

	res = isSubsequence("", "ahbgdc")

	if !res {
		t.Fail()
	}

	res = isSubsequence("abc", "")

	if res {
		t.Fail()
	}

	res = isSubsequence("ahbgdc", "abc")

	if res {
		t.Fail()
	}

	res = isSubsequence("abc", "ahbgdc")

	if !res {
		t.Fail()
	}

	res = isSubsequence("axc", "ahbgdc")

	if res {
		t.Fail()
	}
}

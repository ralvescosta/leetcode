package main

import "testing"

func TestStrStr(t *testing.T) {
	res := StrStr("mississippi", "issip")

	if res != 4 {
		t.Error("error")
	}
}

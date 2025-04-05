package main

import (
	"testing"
)

func Test_isAnagram(t *testing.T) {
	anagramTests := []struct {
		name        string
		testString1 string
		testString2 string
		expected    bool
		msg         string
	}{
		{"valid anagram", "zoo", "ooz", true, "ooz is an anagram of zoo"},
		{"must use the same letters", "abc", "def", false, "def is not an anagram of abc"},
		{"must have the same word length", "abc", "dcba", false, "dcba is not an anagram of abc"},
		{"must not be the same words", "abc", "abc", false, "abc is not an anagram of abc"},
	}

	for _, e := range anagramTests {
		result, msg := isAnagram(e.testString1, e.testString2)
		if e.expected && !result {
			t.Errorf("%s: expected true but got false", e.name)
		}
		if !e.expected && result {
			t.Errorf("%s: expected false but got true", e.name)
		}

		if e.msg != msg {
			t.Errorf("%s: expected %s but got %s", e.name, e.msg, msg)
		}
	}

}

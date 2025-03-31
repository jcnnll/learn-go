package main

import "fmt"

func main() {
	_, msg := isAnagram("abba", "abba")
	fmt.Println(msg)

	_, msg = isAnagram("zoo", "ozo")
	fmt.Println(msg)

	_, msg = isAnagram("abc", "def")
	fmt.Println(msg)
}

func isAnagram(s1 string, s2 string) (bool, string) {
	if s1 == s2 {
		return false, fmt.Sprintf("%v is not an anagram of %v", s2, s1)
	}

	if len(s1) != len(s2) {
		return false, fmt.Sprintf("%v is not an anagram of %v", s2, s1)
	}

	m := make(map[rune]int)

	for _, c := range s1 {
		_, ok := m[c]
		if ok {
			m[c] += 1
			continue
		}
		m[c] = 1
	}

	for _, c := range s2 {
		_, ok := m[c]
		if ok {
			m[c] -= 1
			continue
		}
		return false, fmt.Sprintf("%v is not an anagram of %v", s2, s1)
	}

	return true, fmt.Sprintf("%v is an anagram of %v", s2, s1)
}

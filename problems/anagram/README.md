# Anagram problem
Given 2 strings, determine whether the second string is an anagram of the first.

## Concrete examples
- `ooz` is an anagram of `zoo`
- `abba` is not an anagram, it's a palindrome
- `abc` is not an anagram of `def`

## Rules
- A valid anagram has the same letters as the original word in a diffrent order
- The anagram must have the same letters as the original word
- The oder of letters may not be the same as the original word letter order
- Must not be the same as the original word

The following test cases will determine whether the above rules are applied:

```
		{"valid anagram", "zoo", "ooz", true, "ooz is an anagram of zoo"},
		{"must use the same letters", "abc", "def", false, "def is not an anagram of abc"},
		{"must have the same word length", "abc", "dcba", false, "dcba is not an anagram of abc"},
		{"must not be the same words", "abc", "abc", false, "abc is not an anagram of abc"},
```

## Solution
- Check that the words have the same length - if not, the second word cannot be an anagram
- Check that the words are not equal to each other - if they are, then the second word cannot be an anagram
- Next check each letter in the first word exists in the second word - thic could be done using a nested loop
but that would be O(n^2) in time complexity - this can be reduced to O(n) by using a frequency counter object
and 2 loops where the first loop adds each occurrence of a letter as a key and count to a map. The secod loop
through the second word checks if the key exists - if not we short circuit the loop and return false otherwise
continue decrementing the count of each letter in the map - if no errors occur, we have a valid anagram
```go

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
```

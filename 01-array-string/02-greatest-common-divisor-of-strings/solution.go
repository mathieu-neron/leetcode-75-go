package solution

import "strings"

// LeetCode 1071: Greatest Common Divisor of Strings
// https://leetcode.com/problems/greatest-common-divisor-of-strings/

func gcdOfStrings(str1, str2 string) string {
	if len(str1) < len(str2) {
		gcdOfStrings(str2, str1)
	}

	if !strings.HasPrefix(str1, str2) {
		return ""
	}
	if len(str2) == 0 {
		return str1
	}

	return gcdOfStrings(str1[len(str2):], str2)
}

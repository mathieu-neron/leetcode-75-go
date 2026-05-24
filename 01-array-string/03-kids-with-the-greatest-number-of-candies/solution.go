package solution

import "slices"

// LeetCode 1431: Kids With the Greatest Number of Candies
// https://leetcode.com/problems/kids-with-the-greatest-number-of-candies/

func kidsWithCandies(candies []int, extraCandies int) []bool {
	n := len(candies)
	mx := slices.Max(candies)
	result := make([]bool, n)

	for i, candy := range candies {
		if candy+extraCandies >= mx {
			result[i] = true
		} else {
			result[i] = false
		}
	}

	return result
}

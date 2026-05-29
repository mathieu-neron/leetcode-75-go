package solution

// LeetCode 605: Can Place Flowers
// https://leetcode.com/problems/can-place-flowers/

func canPlaceFlowers(flowerbed []int, n int) bool {
	f := len(flowerbed)
	i := 0

	for i < f {
		if flowerbed[i] == 1 {
			i += 2
		} else if i == f-1 || flowerbed[i+1] == 0 {
			n--
			i = i + 2
		} else {
			i += 3
		}
	}
	return n <= 0
}

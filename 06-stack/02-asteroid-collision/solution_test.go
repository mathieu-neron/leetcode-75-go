package solution

import (
	"testing"
	"reflect"
)

func TestAsteroidCollision(t *testing.T) {
	t.Skip("TODO: implement asteroidCollision")
	tests := []struct {
		asteroids []int
		want []int
	}{
		{[]int{5, 10, -5}, []int{5, 10}},
		{[]int{8, -8}, []int{}},
		{[]int{10, 2, -5}, []int{10}},
		{[]int{-2, -1, 1, 2}, []int{-2, -1, 1, 2}},
	}
	for _, tc := range tests {
		got := asteroidCollision(tc.asteroids)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("asteroidCollision(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}

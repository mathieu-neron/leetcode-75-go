package solution

import (
	"testing"
	"reflect"
)

func TestDailyTemperatures(t *testing.T) {
	t.Skip("TODO: implement dailyTemperatures")
	tests := []struct {
		temperatures []int
		want []int
	}{
		{[]int{73, 74, 75, 71, 69, 72, 76, 73}, []int{1, 1, 4, 2, 1, 1, 0, 0}},
		{[]int{30, 40, 50, 60}, []int{1, 1, 1, 0}},
		{[]int{30, 60, 90}, []int{1, 1, 0}},
	}
	for _, tc := range tests {
		got := dailyTemperatures(tc.temperatures)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("dailyTemperatures(%v) = %v, want %v", tc, got, tc.want)
		}
	}
}

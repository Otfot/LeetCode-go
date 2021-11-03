package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSumII(t *testing.T) {
	type param struct {
		numbers []int
		target  int
	}

	type answer struct {
		one []int
	}

	data := []struct {
		param
		answer
	}{
		{
			param{[]int{2, 7, 11, 15}, 9},
			answer{[]int{1, 2}},
		},
		{
			param{[]int{2, 3, 7, 7, 9}, 16},
			answer{[]int{3, 5}},
		},
		{
			param{[]int{2, 6, 9, 12, 14, 18}, 3},
			answer{},
		},
	}

	for _, pa := range data {
		actual := twoSum(pa.param.numbers, pa.param.target)

		if !reflect.DeepEqual(pa.answer.one, actual) {
			t.Errorf("twoSum(%v, %v) = %v; expected %v", pa.param.numbers, pa.param.target, actual, pa.answer.one)
		}
	}
}

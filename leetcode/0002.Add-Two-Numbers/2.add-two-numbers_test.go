package leetcode

import (
	"reflect"
	"testing"
	. "github.com/otfot/LeetCode-go/tools"
)

func TestTwoSum(t *testing.T) {
	type param struct {
		l1  []int
		l2  []int
	}

	type answer struct {
		one []int
	}

	data := []struct {
		param
		answer
	}{
		{
			param{},
			answer{},
		},

		{
			param{[]int{1}, []int{1}},
			answer{[]int{2}},
		},

		{
			param{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
			answer{[]int{2, 4, 6, 8}},
		},

		{
			param{[]int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
			answer{[]int{2, 4, 6, 8, 0, 1}},
		},

		{
			param{[]int{1}, []int{9, 9, 9, 9, 9}},
			answer{[]int{0, 0, 0, 0, 0, 1}},
		},

		{
			param{[]int{9, 9, 9, 9, 9}, []int{1}},
			answer{[]int{0, 0, 0, 0, 0, 1}},
		},

		{
			param{[]int{2, 4, 3}, []int{5, 6, 4}},
			answer{[]int{7, 0, 8}},
		},

		{
			param{[]int{1, 8, 3}, []int{7, 1}},
			answer{[]int{8, 9, 3}},
		},
	}

	for _, pa := range data {
		actual := addTwoNumbers(Ints2NewListNode(pa.param.l1), Ints2NewListNode(pa.param.l2))

		arrActual := NewListNode2Ints(actual)
		
		 if !reflect.DeepEqual(arrActual, pa.answer.one) {
			t.Errorf("twoSum(%v, %v) = %v; expected %v", pa.param.l1, pa.param.l2, arrActual, pa.answer.one)
		}
	}

}

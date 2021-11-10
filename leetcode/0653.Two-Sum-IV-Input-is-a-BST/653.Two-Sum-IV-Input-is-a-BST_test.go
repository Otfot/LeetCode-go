package leetcode

import (
	"reflect"
	"testing"

	. "github.com/otfot/LeetCode-go/tools"
)

func TestFindTarget(t *testing.T) {
	type param struct {
		root []int
		k    int
	}

	type answer struct {
		one bool
	}

	data := []struct {
		param
		answer
	}{
		{
			param{[]int{}, 3},
			answer{false},
		},
		{
			param{[]int{3, 9, 20, NULL, NULL, 15, 7}, 29},
			answer{true},
		},

		{
			param{[]int{1, 2, 3, 4, NULL, NULL, 5}, 9},
			answer{true},
		},

		{
			param{[]int{1, 2, 3, 4, NULL, 5}, 4},
			answer{true},
		},
	}

	for _, pa := range data {
		actual := findTarget(Ints2NewTreeNode(pa.param.root), pa.param.k)

		if !reflect.DeepEqual(actual, pa.answer.one) {
			t.Errorf("findTarget(%v, %v) = %v; expected %v", Ints2NewTreeNode(pa.param.root), pa.param.k, actual, pa.answer.one)
		}
	}

}

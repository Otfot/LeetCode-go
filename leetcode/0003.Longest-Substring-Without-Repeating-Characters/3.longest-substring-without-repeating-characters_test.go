package leetcode

import (
	"reflect"
	"testing"
)

func TestTwoSum(t *testing.T) {
	type param struct {
		s string
	}

	type answer struct {
		one int
	}

	data := []struct {
		param
		answer
	}{
		{
			param{"awedkwd"},
			answer{5},
		},
		{
			param{"wwwwwww"},
			answer{1},
		},
	}

	for _, pa := range data {
		actual := lengthOfLongestSubstring(pa.param.s)

		if !reflect.DeepEqual(actual, pa.answer.one) {
			t.Errorf("twoSum(%v) = %v; expected %v", pa.param.s, actual, pa.answer.one)
		}
	}

}

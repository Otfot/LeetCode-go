package leetcode

import (
	"reflect"
	"testing"
)

func TestLongestPalindrome(t *testing.T) {
	type param struct {
		s string
	}

	type answer struct {
		one string
	}

	data := []struct {
		param
		answer
	}{
		{
			param{"awedkwd"},
			answer{"a"},
		},
		{
			param{"wwwwwww"},
			answer{"wwwwwww"},
		},
		{
			param{"wwwabwww"},
			answer{"www"},
		},
	}

	for _, pa := range data {
		actual := longestPalindrome(pa.param.s)

		if !reflect.DeepEqual(actual, pa.answer.one) {
			t.Errorf("longestPalindrome(%v) = %v; expected %v", pa.param.s, actual, pa.answer.one)
		}
	}

}

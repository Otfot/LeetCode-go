package leetcode

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
	type param struct {
		s       string
		numRows int
	}

	type answer struct {
		one string
	}

	data := []struct {
		param
		answer
	}{
		{
			param{"PAYPALISHIRING", 3},
			answer{"PAHNAPLSIIGYIR"},
		},
		{
			param{"w", 3},
			answer{"w"},
		},
		{
			param{"PAYPALISHIRING", 4},
			answer{"PINALSIGYAHRPI"},
		},
	}

	for _, pa := range data {
		actual := convert(pa.param.s, pa.param.numRows)

		if !reflect.DeepEqual(actual, pa.answer.one) {
			t.Errorf("convert(%v, %v) = %v; expected %v", pa.param.s, pa.param.numRows, actual, pa.answer.one)
		}
	}

}

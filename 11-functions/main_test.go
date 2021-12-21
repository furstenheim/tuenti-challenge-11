package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCommonPrefix (t *testing.T) {
	tcs := []struct{
		a, b string
		expected int
	} {
		{`a`, `ab`, 1},
		{`ab`, `a`, 1},
		{`dfafacadc`, `eacacna`, 0},
		{`dfafacadc`, `dfeacacna`, 2},
	}

	for _, tc := range tcs {
		res := findCommonRoot(tc.a, tc.b)
		assert.Equal(t, tc.expected, res)
	}
}

func TestBestSolution (t *testing.T) {
	tcs := []struct{
		a Case
		expected string
	} {
		{
			a:        Case{
				NTotalFunctions:   2,
				NFunctionsPerFile: 1,
				Functions:         []string{"a", "b"},
			},
			expected: "2",
		},
		{
			a:        Case{
				NTotalFunctions:   4,
				NFunctionsPerFile: 2,
				Functions:         []string{"a", "ab", "abc", "ad"},
			},
			expected: "3",
		},
	}

	for _, tc := range tcs {
		res := solveCase(tc.a)
		assert.Equal(t, tc.expected, res)
	}
}

package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFactorizeNumber(t *testing.T) {
	tcs := []struct {
		c        int
		expected []PrimePower
	}{
		{
			c:        19,
			expected: []PrimePower{{
				base:  19,
				power: 1,
			}},
		},
		{
			c:        8,
			expected: []PrimePower{{
				base:  2,
				power: 3,
			}},
		},
		{
			c:        6,
			expected: []PrimePower{{
				base:  2,
				power: 1,
			},{
				base:  3,
				power: 1,
			}},
		},
	}

	for _, tc := range tcs {
		res := FactorizeNumber(tc.c)
		assert.Equal(t, tc.expected, res)
	}
}


func TestMerge(t *testing.T) {
	tcs := []struct {
		c1        []PrimePower
		c2        []PrimePower
		expected []PrimePower
	}{
		{

			c1: []PrimePower{{
				base:  19,
				power: 1,
			}},
			c2: []PrimePower{{
				base:  19,
				power: 1,
			}},
			expected: []PrimePower{{
				base:  19,
				power: 1,
			}},
		},{

			c1: []PrimePower{{
				base:  19,
				power: 1,
			}},
			c2: []PrimePower{{
				base:  3,
				power: 1,
			}},
			expected: []PrimePower{{
				base:  19,
				power: 1,
			},
			{
				base:  3,
				power: 1,
			}},
		},{

			c1: []PrimePower{{
				base:  3,
				power: 1,
			}},
			c2: []PrimePower{{
				base:  3,
				power: 2,
			}},
			expected: []PrimePower{{
				base:  3,
				power: 2,
			}},
		},
	}

	for _, tc := range tcs {
		res := mergeDivisors(tc.c1, tc.c2)
		assert.Equal(t, tc.expected, res)
	}
}


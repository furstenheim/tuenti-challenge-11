package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGraph(t *testing.T) {
	tcs := []struct {
		c        Case
		expected map[CryptoId]map[CryptoId]int
	}{
		{
			c: Case{
				NMarkets: 2,
				Markets: []Market{{
					Name:       "",
					NExchanges: 2,
					Exchanges: []Exchange{{
						Start:  "A",
						Amount: 1,
						End:    "B",
					}, {
						Start:  "B",
						Amount: 2,
						End:    "C",
					}},
				},
					{
						Name:       "",
						NExchanges: 2,
						Exchanges: []Exchange{{
							Start:  "A",
							Amount: 2,
							End:    "B",
						}, {
							Start:  "B",
							Amount: 1,
							End:    "C",
						}},
					},
				},
			},
			expected: map[CryptoId]map[CryptoId]int{
				"A": {"B": 2},
				"B": {"C": 2},
			},
		},
		{
			c: Case{
				NMarkets: 2,
				Markets: []Market{{
					Name:       "",
					NExchanges: 2,
					Exchanges: []Exchange{{
						Start:  "A",
						Amount: 1,
						End:    "B",
					}, {
						Start:  "B",
						Amount: 2,
						End:    "C",
					}, {
						// Several starting from same letter
						Start:  "A",
						Amount: 3,
						End:    "D",
					}},
				},
					{
						Name:       "",
						NExchanges: 2,
						Exchanges: []Exchange{{
							Start:  "A",
							Amount: 2,
							End:    "B",
						}, {
							Start:  "B",
							Amount: 1,
							End:    "C",
						}},
					},
				},
			},
			expected: map[CryptoId]map[CryptoId]int{
				"A": {"B": 2, "D": 3},
				"B": {"C": 2},
			},
		},
	}

	for _, tc := range tcs {
		res := computeGraph(tc.c)
		assert.Equal(t, tc.expected, res)
	}
}


func TestCalculation(t *testing.T) {
	tcs := []struct {
		c        Case
		expected string
	}{
		{
			// First path is not the fastest
			c: Case{
				NMarkets: 2,
				Markets: []Market{{
					Name:       "",
					NExchanges: 2,
					Exchanges: []Exchange{{
						Start:  "BTC",
						Amount: 1,
						End:    "A",
					}, {
						Start:  "A",
						Amount: 2,
						End:    "C",
					}, {
						Start:  "A",
						Amount: 1,
						End:    "D",
					}},
				},
					{
						Name:       "",
						NExchanges: 2,
						Exchanges: []Exchange{{
							Start:  "C",
							Amount: 1,
							End:    "A",
						}, {
							Start:  "D",
							Amount: 3,
							End:    "A",
						}, {
							Start:  "A",
							Amount: 1,
							End:    "BTC",
						}},
					},
				},
			},
			expected: "3",
		},{
			// Fastest path has lower quantity
			c: Case{
				NMarkets: 2,
				Markets: []Market{{
					Name:       "",
					NExchanges: 2,
					Exchanges: []Exchange{{
						Start:  "BTC",
						Amount: 1,
						End:    "A",
					}, {
						Start:  "A",
						Amount: 2,
						End:    "C",
					}, {
						Start:  "A",
						Amount: 1,
						End:    "D",
					}},
				},
					{
						Name:       "",
						NExchanges: 2,
						Exchanges: []Exchange{{
							Start:  "C",
							Amount: 1,
							End:    "A",
						}, {
							Start:  "D",
							Amount: 3,
							End:    "A",
						}, {
							Start:  "A",
							Amount: 2,
							End:    "BTC",
						}},
					},
				},
			},
			expected: "2",
		},
	}

	for _, tc := range tcs {
		res := solveCase(tc.c)
		assert.Equal(t, tc.expected, res)
	}
}

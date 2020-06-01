package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestPopulation_Fitness(t *testing.T) {
	cond := Condition{1, 2, 3, 4, 30}
	p := Population{
		bodies: []Body{
			{
				c:     cond,
				genes: [4]int{1, 28, 15, 3},
			},
			{
				c:     cond,
				genes: [4]int{14, 9, 2, 4},
			},
			{
				c:     cond,
				genes: [4]int{13, 5, 7, 3},
			},
			{
				c:     cond,
				genes: [4]int{23, 8, 16, 19},
			},
			{
				c:     cond,
				genes: [4]int{9, 13, 5, 2},
			},
		},
	}

	require.Equal(t, float32(59), p.Fitness())
}

func TestPopulation_pairs(t *testing.T) {
	p := Population{
		bodies: []Body{
			{},
			{},
			{},
			{},
			{},
		},
	}
	require.Equal(t, [][2]int{
		{0, 1}, {0, 2}, {0, 3}, {0, 4},
		{1, 0}, {1, 2}, {1, 3}, {1, 4},
		{2, 0}, {2, 1}, {2, 3}, {2, 4},
		{3, 0}, {3, 1}, {3, 2}, {3, 4},
		{4, 0}, {4, 1}, {4, 2}, {4, 3},
	}, p.pairs())
}

func TestPopulation_Reproduction(t *testing.T) {
	cond := Condition{1, 2, 3, 4, 30}
	p1 := makePopulation(cond, 5)
	require.Equal(t, 5, p1.Count())

	p2 := p1.Reproduction()
	require.Equal(t, 20, p2.Count())
}

func TestPopulation_Select(t *testing.T) {
	cond := Condition{1, 1, 1, 1, 4}
	p1 := Population{
		bodies: []Body{
			{
				c:     cond,
				genes: [4]int{2, 2, 2, 2},
			},
			{
				c:     cond,
				genes: [4]int{4, 4, 4, 4},
			},
			{
				c:     cond,
				genes: [4]int{1, 1, 1, 1},
			},
			{
				c:     cond,
				genes: [4]int{3, 3, 3, 3},
			},
			{
				c:     cond,
				genes: [4]int{5, 5, 5, 5},
			},
		},
	}

	p2 := p1.Select(3)

	require.Equal(t, 3, p2.Count())
	require.Equal(t, Population{
		bodies: []Body{
			{
				c:     cond,
				genes: [4]int{1, 1, 1, 1},
			},
			{
				c:     cond,
				genes: [4]int{2, 2, 2, 2},
			},
			{
				c:     cond,
				genes: [4]int{3, 3, 3, 3},
			},
		},
	}, p2)
}

func TestPopulation_Adapted(t *testing.T) {
	cond := Condition{1, 1, 1, 1, 10}
	p := Population{
		bodies: []Body{
			{
				// приспособился
				c:     cond,
				genes: [4]int{10, 0, 0, 0},
			},
			{
				c:     cond,
				genes: [4]int{1, 1, 1, 1},
			},
			{
				c:     cond,
				genes: [4]int{3, 3, 3, 3},
			},
			{
				c:     cond,
				genes: [4]int{5, 5, 5, 5},
			},
			{
				// приспособился
				c:     cond,
				genes: [4]int{1, 1, 4, 4},
			},
		},
	}

	require.Equal(t, []Body{
		{
			c:     cond,
			genes: [4]int{10, 0, 0, 0},
		},
		{
			c:     cond,
			genes: [4]int{1, 1, 4, 4},
		},
	}, p.Adapted())
}

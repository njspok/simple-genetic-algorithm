package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBody_Prox(t *testing.T) {
	b := Body{
		c:     Condition{1, 2, 3, 4, 30},
		genes: [4]int{1, 28, 15, 3},
	}

	require.Equal(t, 84, b.Prox())
}

func TestBody_cross(t *testing.T) {
	cond := Condition{1, 2, 3, 4, 30}
	b1 := Body{
		c:     cond,
		genes: [4]int{11, 22, 33, 44},
	}
	b2 := Body{
		c:     cond,
		genes: [4]int{55, 66, 77, 88},
	}

	require.Equal(t, Body{
		c:     cond,
		genes: [4]int{11, 66, 77, 88},
	}, b1.cross(1, b2))
	require.Equal(t, Body{
		c:     cond,
		genes: [4]int{11, 22, 77, 88},
	}, b1.cross(2, b2))
	require.Equal(t, Body{
		c:     cond,
		genes: [4]int{11, 22, 33, 88},
	}, b1.cross(3, b2))
}

func TestBody_mutate(t *testing.T) {
	originGenes := [4]int{11, 22, 33, 44}

	// create body from genes
	b := Body{c: Condition{1, 2, 3, 4, 30}}
	copy(b.genes[:], originGenes[:])

	// process it!
	b.mutate()

	// check, only 1 difference
	diff := 0
	for i, _ := range originGenes {
		if originGenes[i] != b.genes[i] {
			diff++
		}
	}
	require.Equal(t, 1, diff)
}

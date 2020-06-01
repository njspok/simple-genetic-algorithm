package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEvolution(t *testing.T) {
	cond := Condition{1, 2, 3, 4, 30}
	bodies, c := Evolution(cond)
	require.NotEmpty(t, bodies)
	require.True(t, c >= 1)
	for _, g := range bodies {
		g := g.Genes()
		require.Equal(t, 30, 1*g[0]+2*g[1]+3*g[2]+4*g[3])
	}
}

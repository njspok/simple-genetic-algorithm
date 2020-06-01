package domain

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCondition_E(t *testing.T) {
	c := Condition{
		11,
		22,
		33,
		44,
		55,
	}
	require.Equal(t, 55, c.E())
}

package domain

// Организм.
// Гены представляю собой значения (x,y,z,w) переменных.
// Организм умеет скрещиваться с другими организмами,
// мутировать гены.

import "math/rand"

type Body struct {
	c     Condition
	genes [4]int
}

func (b Body) Genes() [4]int {
	return b.genes
}

func (b Body) IsAdapted() bool {
	return b.Prox() == 0
}

func (b Body) Prox() int {
	res := 0
	for i, v := range b.genes {
		res += v * b.c[i]
	}
	return res - b.c.E()
}

func (b Body) Cross(b2 Body) Body {
	n := rand.Intn(len(b.genes)) + 1
	child := b.cross(n, b2)
	child.mutate()
	return child
}

func (b *Body) mutate() {
	i := rand.Intn(len(b.genes))
	v := rand.Intn(b.c.E()) + 1
	b.genes[i] = v
}

func (b Body) cross(n int, b2 Body) Body {
	child := Body{c: b.c}
	g1 := b.genes[0:n]
	g2 := b2.genes[n:]
	newGenes := append(g1, g2...)
	copy(child.genes[:], newGenes)
	return child
}

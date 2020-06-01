package domain

// Эволюция.
// Собственно поиск решения уравнения, путем порождения
// новых поколей (новых решений) и отбора наиболее приспособившихся
// особей (лучших решений).

import (
	"math/rand"
	"time"
)

func init() {
	// todo не очень хорошо
	rand.Seed(time.Now().UnixNano())
}

func makeBody(c Condition) Body {
	return Body{
		c: c,
		genes: [4]int{
			rand.Intn(c.E() + 1),
			rand.Intn(c.E() + 1),
			rand.Intn(c.E() + 1),
			rand.Intn(c.E() + 1),
		},
	}
}

func makePopulation(c Condition, count int) Population {
	pop := Population{}
	for i := 0; i < count; i++ {
		pop.bodies = append(pop.bodies, makeBody(c))
	}
	return pop
}

func Evolution(c Condition) ([]Body, int) {
	genCounter := 0

	pop := makePopulation(c, 50)
	if bodies := pop.Adapted(); len(bodies) > 0 {
		return bodies, genCounter
	}

	var next Population
	for {
		genCounter++
		next = pop.Reproduction()
		if bodies := next.Adapted(); len(bodies) > 0 {
			return bodies, genCounter
		}
		next = next.Select(50)
	}
}

package domain

// Эволюция.
// Собственно поиск решения уравнения, путем порождения
// новых поколей (новых решений) и отбора наиболее приспособившихся
// особей (лучших решений).

import (
	"fmt"
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

	// генерируем случайную популяцию и вдруг нам повезет
	old := makePopulation(c, 5)
	fmt.Println("iteration 0: fitness", old.Fitness())
	if bodies := old.Adapted(); len(bodies) > 0 {
		return bodies, genCounter
	}

	var next Population
	for {
		genCounter++

		// получаем новое поколение и вдруг нам повезет
		next = old.Reproduction()
		if bodies := next.Adapted(); len(bodies) > 0 {
			next = next.Select(5) // отбираем для правильного вычисления Fitness
			fmt.Printf("iteration %v: fitness %v prox %v\n", genCounter, next.Fitness(), prox(next))
			return bodies, genCounter
		}

		// отбираем лучших
		next = next.Select(5)

		fmt.Printf("iteration %v: fitness %v to %v prox %v\n", genCounter, old.Fitness(), next.Fitness(), prox(next))

		// если новое поколение лучше старого, то берем его,
		// в противном случае - отказывается от него и пробуем
		// еще раз со старым поколением
		if next.Fitness() <= old.Fitness() {
			old = next
		}
	}
}

// вернет массив
func prox(p Population) []int {
	var result []int
	for _, b := range p.Bodies() {
		result = append(result, b.Prox())
	}
	return result
}

package domain

// Популяция.
// Сборище организмов. В популяции можно найти организмы,
// которые идеально приспособились к условиям среды (это и есть
// решения уравнения), выбрать наиболее приспособившихся особей
// для генерации следующего поколения.

import "sort"

type Population struct {
	bodies []Body
}

// Bodies возвращаем особей популяции
func (p Population) Bodies() []Body {
	return p.bodies
}

// Reproduction создаем новую популяцию, на основе
// старой, путем скрещивания каждого с каждым.
func (p Population) Reproduction() Population {
	pairs := p.pairs()
	newPop := Population{}
	for _, pair := range pairs {
		b1 := p.bodies[pair[0]]
		b2 := p.bodies[pair[1]]

		child := b1.Cross(b2)

		newPop.bodies = append(newPop.bodies, child)
	}
	return newPop
}

func (p Population) pairs() [][2]int {
	var res [][2]int
	for i := 0; i < len(p.bodies); i++ {
		for j := 0; j < len(p.bodies); j++ {
			if i != j {
				res = append(res, [2]int{i, j})
			}
		}
	}
	return res
}

func (p Population) Count() int {
	return len(p.bodies)
}

// Fitness средний показатель приспособленности
// особей популяции.
func (p Population) Fitness() float32 {
	var res float32 = 0
	for _, b := range p.bodies {
		res += float32(b.Prox())
	}
	return res / float32(p.Count())
}

// Select создает новую популяции, на основании отбора n
// лучших особей из прежней популяции.
func (p Population) Select(n int) Population {
	var proxes [][2]int
	for i, b := range p.bodies {
		proxes = append(proxes, [2]int{i, b.Prox()})
	}

	sort.SliceStable(proxes, func(i, j int) bool {
		return proxes[i][1] < proxes[j][1]
	})

	// truncate
	proxes = proxes[:n]

	var bodies []Body
	for _, pr := range proxes {
		bodies = append(bodies, p.bodies[pr[0]])
	}

	return Population{
		bodies: bodies,
	}
}

// Adapted вернет список особей, которые адаптировались к условиям.
func (p Population) Adapted() []Body {
	var res []Body
	for _, b := range p.bodies {
		if b.IsAdapted() {
			res = append(res, b)
		}
	}
	return res
}

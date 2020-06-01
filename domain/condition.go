package domain

// Условия среды обитания организма, под которые он подстраивается,
// те самые параметры (A,B,C,D,E) уравнения.

type Condition [5]int

func (c Condition) factor() [4]int {
	return [4]int{
		c[0], // A
		c[1], // B
		c[2], // C
		c[3], // D
	}
}

func (c Condition) E() int {
	return c[4]
}

package main

import (
	"fmt"
	"github.com/njspok/simple-genetic-algorithm/domain"
)

func main() {
	cond := domain.Condition{1, 2, 3, 4, 30}
	bodies, generations := domain.Evolution(cond)
	fmt.Printf("%v×x + %v×y + %v×z + %v×w = %v\n", cond[0], cond[1], cond[2], cond[3], cond[4])
	for i, b := range bodies {
		fmt.Printf("Solve %v: %v\n", i+1, b.Genes())
	}
	fmt.Println("Generations: ", generations)
}

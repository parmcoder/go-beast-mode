package main

import (
	"errors"
	"fmt"
)

const limA = 2000

type EulerSolution struct {
	first   int
	second  int
	third   int
	fourth  int
	culprit int
}

func intPow(n, m int) int {
	if m == 0 {
		return 1
	}
	result := n
	for i := 2; i <= m; i++ {
		result *= n
	}
	return result
}

func computeEuler() (EulerSolution, error) {
	for e := 1; e < limA; e++ {
		for a := 1; a < e; a++ {
			var a5 = intPow(a, 5)
			for b := 1; b <= a; b++ {
				var b5 = intPow(b, 5)
				for c := 1; c < b; c++ {
					var c5 = intPow(c, 5)
					for d := 1; d < c; d++ {
						var d5 = intPow(d, 5)
						var e5 = intPow(e, 5)
						var got int = a5 + b5 + c5 + d5

						if got == e5 {
							return EulerSolution{
								first:   a,
								second:  b,
								third:   c,
								fourth:  d,
								culprit: e,
							}, nil
						}

					}
				}
			}
		}
	}

	return EulerSolution{}, errors.New("Euler solution")
}

func computeBeastEuler() (EulerSolution, error) {
	for e := 1; e < limA; e++ {
		var e5 = intPow(e, 5)
		for a := 1; a < e; a++ {
			var a5 = intPow(a, 5)
			for b := 1; b <= a; b++ {
				var b5 = intPow(b, 5)
				for c := 1; c < b; c++ {
					var c5 = intPow(c, 5)
					for d := 1; d < c; d++ {
						var d5 = intPow(d, 5)
						var got int = a5 + b5 + c5 + d5
						if got == e5 {
							return EulerSolution{
								first:   a,
								second:  b,
								third:   c,
								fourth:  d,
								culprit: e,
							}, nil
						}

					}
				}
			}
		}
	}

	return EulerSolution{}, errors.New("Euler solution")
}

func main() {
	solution, err := computeEuler()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(solution)
}

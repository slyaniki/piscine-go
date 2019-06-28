package piscine

import "math"

func IsPrime(nb int) bool {
	for i := 2; i <= int(math.Floor(float64(nb) / 2)); i++ {
		if nb%i == 0 {
			return false
		}
	}
	return nb > 1
}
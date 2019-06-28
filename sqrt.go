package piscine

import "math"

func Sqrt(nb int) int {
	aux:=math.Sqrt(float64(nb))
	aux2 :=int(aux)
	if float64(aux2) == aux{
		return int(aux)
	} else{
		return 0
	}
}
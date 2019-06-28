package piscine

import (
    "math"
    "fmt"
)
func FindNextPrime(nb int) int {
	for !IsPrime(nb){
		nb+=1
	}
	return nb
}

type Point struct {
	x int
	y int
}

var results = make([][]Point, 0)


func EightQueens() {
	Solve(8)

}

func Solve(n int) {
	for col := 0; col < n; col++ {
		start := Point{x: col, y: 0}
		current := make([]Point, 0)
		Recurse(start, current, n)
	}
	for _, result := range results {
		for _,point := range result{
			fmt.Print(point.x+1)
		}
		fmt.Print("\n")
	}
}
func Recurse(point Point, current []Point, n int) {
	if CanPlace(point, current) {
		current = append(current, point)
		if len(current) == n {
			c := make([]Point, n)
			for i, point := range current {
				c[i] = point
			}
			results = append(results, c)
		} else {
			for col := 0; col < n; col++ {
				for row := point.y; row < n; row++ {
					nextStart := Point{x: col, y: row}
					Recurse(nextStart, current, n)

				}

			}
		}
	}
}
func CanPlace(target Point, board []Point) bool {
	for _, point := range board {
		if CanAttack(point, target) {
			return false
		}
	}
	return true
}

func CanAttack(a, b Point) bool {
	//fmt.Print(a, b)
	answer := a.x == b.x || a.y == b.y || math.Abs(float64(a.y-b.y)) == math.Abs(float64(a.x-b.x))
	//fmt.Print(answer)
	return answer
}
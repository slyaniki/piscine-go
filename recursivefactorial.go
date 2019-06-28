package piscine
 

func RecursiveFactorial(nb int) int{
	if nb>16 || nb<0{
		return 0
	} else if nb <=1{
		return 1
	} else{
		return nb * RecursiveFactorial(nb -1)
	}
}

package piscine


func IterativeFactorial(nb int) int{
	if nb>16 || nb<0{
		return 0
	} else {
		s:=1
		for i:=1; i<=nb; i++{
			s*= i
			if s>2147483647 || s<0 {
				return 0
			}
		}
		return s
	}

}
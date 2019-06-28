package piscine


func IterativePower(nb, power int) int{
	if power<0{
		return 0
	}else if power==0{
		return 1
	}  else {
		p := 1
		for i := 1; i <= power; i++ {
			p *= nb
		}
		if p>2147483647{
			return 0
		} else{
			return p
		}
	}
}
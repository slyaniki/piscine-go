package piscine

import (
	"fmt"
)

func printlinec( n,x,y int) {
	if(x == 1){
		fmt.Print("A")
		for i:= 0; i<n-2 ;i++ {
			fmt.Print("B")
		}
		if n >=2 {
			fmt.Print("A")
		}
		fmt.Print("\n")
	}else if (x==y){
		fmt.Print("C")
		for i:= 0; i<n-2 ;i++ {
			fmt.Print("B")
		}
		if n >=2 {
			fmt.Print("C")
		}
		fmt.Print("\n")
	}else {
		fmt.Print("B")
		for i:= 0; i<n-2 ;i++ {
			fmt.Print(" ")
		}
		if n >=2 {
			fmt.Print("B")
		}
		fmt.Print("\n")
	}
}

func Raid1c(x,y int)  {
	if x>0 && y>0{
		for i:=1;i<=y;i++{
			printlinec(x,i,y)
		}
	}
}

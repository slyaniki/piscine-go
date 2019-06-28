package piscine

import (
	"fmt"
)

func printlinea( n,x,y int) {
	if(x == 1 || x==y){
		fmt.Print("o")
		for i:= 0; i<n-2 ;i++ {
			fmt.Print("-")
		}
		if n >=2 {
			fmt.Print("o")
		}
		fmt.Print("\n")
	} else {
		fmt.Print("|")
		for i:= 0; i<n-2 ;i++ {
			fmt.Print(" ")
		}
		if n >=2 {
			fmt.Print("|")
		}
		fmt.Print("\n")
	}
}

func Raid1a(x,y int)  {
	if x>0 && y>0{
		for i:=1;i<=y;i++{
			printlinea(x,i,y)
		}
	}
}

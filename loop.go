package main

import "fmt"

func main() {
	i:=1
	for i <=3 {
		fmt.Println(i)
		i=i+1
	}
	for n:=0 ;n<=5;n++{
		if n%2==0{
			continue
		}
		fmt.Println(n)
	}
	if n:=9 ;n<0{
		fmt.Println("xxxx")
	} else {
		fmt.Println("xxsds")
	}

}

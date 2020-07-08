package main

import "fmt"

func main() {
	//与数组不同，slice 的类型仅由它所包含的元素的类型决定
	// 要创建一个长度不为 0 的空 slice，需要使用内建函数 make。
	s:=make([]string,3)
	fmt.Println(s)
	s[0]="s"
	s[1]="ss"
	s[2]="ssss"
	s=append(s,"qqq","sdsd")
	fmt.Println(s)
	twoD:=make([][]int,3)
	for i:=0;i<3;i++{
		innerLen:=i+1
		twoD[i]=make([]int,innerLen)
		for j:=0;j<innerLen;j++{
			twoD[i][j]=i+j
		}

	}
	fmt.Println(twoD)
}

package main

import "fmt"

func main() {
	i:=1
	zeoval(i)
	fmt.Println(i)
	zeopal(&i)//通过 &i 语法来取得 i 的内存地址，即指向 i 的指针。

	fmt.Println(i)

}
/*zeroptr 有一个和上面不同的参数：*int，
这意味着它使用了 int 指针。
紧接着，函数体内的 *iptr 会 解引用 这个指针，
从它的内存地址得到这个地址当前对应的值。
对解引用的指针赋值，会改变这个指针引用的真实地址的值*/
func zeoval(x  int ){
	x=0
}
func zeopal (x *int ){
	*x =0
}
package main

import (
	"errors"
	"fmt"
)

func main() {
	for _, i := range []int{7, 42} {
		if r, e := f1(i); e != nil {
			fmt.Println("f1 failed:", e)
		} else {
			fmt.Println("f1 worked:", r)
		}
	}
	for _, i := range []int{7, 42} {
		if r, e := f2(i); e != nil {
			fmt.Println("f2 failed:", e)
		} else {
			fmt.Println("f2 worked:", r)
		}
	}

	_, e := f2(42)
	if ae, ok := e.(*argError); ok {
		fmt.Println(ae.arg)
		fmt.Println(ae.prob)
	}
}
func f1(x int) (i int, int error) {
	if x==42{
		return -1, errors.New("dddd")

	}
	return x ,nil
}
//你还可以通过实现 Error() 方法来自定义 error 类型。
//这里使用自定义错误类型来表示上面例子中的参数错误。
type argError struct {
	arg  int
	prob string

}
func (e *argError)Error()string{
	return fmt.Sprintf("%d - %s", e.arg, e.prob)
}
func f2(arg int) (int, error) {
	if arg == 42 {

		return -1, &argError{arg, "can't work with it"}
	}
	return arg + 3, nil
}
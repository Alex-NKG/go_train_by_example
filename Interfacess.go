package main

import (
	"fmt"
	"math"
)

func main() {
	r := rect{width: 3, height: 4}
	c := circle{radius: 5}
	measure(r)
	measure(c)

}
// interface 是一种具有一组方法的类型，这些方法定义了 interface 的行为。
type gemo interface {
	area()float64
	perim()float64

}
type rect struct {//定义结构
	width,height float64
}
type circle struct {//定义结构
	radius float64
}
func (r rect) area() float64 {//方法
	return r.width * r.height
}
func (r rect) perim() float64 {//方法
	return 2*r.width + 2*r.height
}
func (c circle) area() float64 {//方法
	return math.Pi * c.radius * c.radius
}
func (c circle) perim() float64 {//方法
	return 2 * math.Pi * c.radius
}
func measure (g gemo){//定义了一个函数 参数类型是 gemo，
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}
/*如果一个变量实现了某个接口，我们就可以调用指定接口中的方法。
这儿有一个通用的 measure 函数，
我们可以通过它来使用所有的 geometry。*/
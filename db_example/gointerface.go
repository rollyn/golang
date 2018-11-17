package main 


import ("fmt")

type Shape interface {
	Area() float64
}

type Rect struct {
    width, height float64
}

func (r Rect) Area() float64 {
    return r.width * r.height
}

func measure(s Shape) {
	fmt.Println(s.Area())
 }

 func main() {
 	r := Rect{width: 3, height: 4}
 	measure(r)
 }

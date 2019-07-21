package main

func main() {

}

type big []int

func Abs(i int) int {
	return i * i
}
func NewBig() big {
	s := []int{2, 3, 5, 7, 11, 13}
	return s
}

func (b big) Len() int {
	return len(b)
}

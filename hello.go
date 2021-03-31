package main

import "fmt"

// func hi(name string) string {
// 	// fmt.Println("hi!" + name)
// 	msg := "hi!" + name
// 	return msg
// }

func hi(name string) (msg string) {
	// fmt.Println("hi!" + name)
	msg = "hi!" + name
	return
}

func main() {
	// var msg string
	// msg = "hello world"
	// var msg = "hello world"
	// const (
	// 	sun = iota
	// 	mon
	// 	tue
	// )
	// fmt.Println(sun, mon, tue)

	// var x int
	// x = 10 % 3
	// x += 3
	// x++
	// fmt.Println(x)

	// var s string
	// s = "hello " + "world"
	// fmt.Println(s)

	// a := true
	// b := false
	// fmt.Println(a && b)
	// fmt.Println(a || b)
	// fmt.Println(!a)

	// a := 5
	// var pa * int
	// pa = &a // &a = aのアドレス
	// // paの領域にあるデータの値 = *pa
	// fmt.Println(pa)
	// fmt.Println(*pa)

	// hi("taguchi")
	fmt.Println(hi("taguchi"))
}
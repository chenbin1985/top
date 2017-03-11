package main

import (
	// "github.com/chenbin/top"
	"fmt"
	"math/cmplx"
	"top"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	// defer print("main func return\n")

	// rand.Seed(time.Now().Unix())
	// fmt.Println("crrent time is ", time.Now())
	// println("random: ", rand.Intn(100))
	// // var x, y = 2, 3
	// x, y := 2, 3
	// x, y = swap(x, y)
	// println("add: 3 + 2 = ", add(3, 2), "swap: 2, 3 => ", x, y)

	// const f string = "%T(%v)\n"
	// fmt.Printf(f, ToBe, ToBe)
	// fmt.Printf(f, MaxInt, MaxInt)
	// fmt.Printf(f, z, z)

	// var (
	// 	i   int
	// 	f32 float32
	// 	b   bool
	// 	s   string
	// )
	// fmt.Printf("%v %v %v %q\n", i, f32, b, s)

	// fmt.Printf("%T %T %T(%v)\n", 1, 1.2, 0.867+0.5i, 0.867+0.5i)

	// i = 1
	// {
	// 	for i := 0; i < 3; i++ {
	// 		print(i, " ")
	// 	}
	// }

	// if i := 4; i > 0 {
	// 	print("\nin if, i:", i)
	// }

	// fmt.Println("\nouter i: ", i)

	// switch j, i := 3, 0; i {
	// case 1:
	// 	println("case1 ", j)
	// case 5:
	// 	print("case2\n")
	// default:
	// 	print("default\n")
	// }

	// type rect struct {
	// 	x int
	// 	y int
	// }

	// k := rect{1, 2}
	// println("x:", k.x, "y:", k.y)
	// k = rect{y: 1}
	// println("x:", k.x, "y:", k.y)
	// kp := &rect{x: 1}
	// println("x:", kp.x, "y:", kp.y)

	// var arr [2]int
	// var slice = []int{1, 2, 3, 4, 5, 6}
	// ms := make([]int, 2, 5)
	// println("arr", arr[0], arr[1], len(arr), cap(arr))
	// println("slice", slice[0], slice[1], len(slice), cap(slice))
	// emptySlice := ms[1:1]
	// println("emptySlice", len(emptySlice), cap(emptySlice))

	// //ms[1] = 1
	// // ms[4] = 4 // runtime error: index out of range
	// println("make slice", ms[0], ms[1], len(ms), cap(ms))

	// newS := slice[1:4]
	// println("new slice", newS[0], newS[1], newS[2], len(newS), cap(newS))

	// ns := append(newS, 8)
	// for i, v := range ns {
	// 	println(i, ":", v)
	// }

	// fmt.Printf("now slice: %v\n", slice)

	// topSlice, err := (*TopData)(&slice).Top(3)
	// if err != nil {
	// 	fmt.Println(err)
	// } else {
	// 	fmt.Printf("Top 3: %v\n", topSlice)
	// 	fmt.Printf("now slice: %v\n", slice)
	// }

	fmt.Println("Seq Top Test...")
	var data top.TopData
	data.New(10 - 0).Seq().ShowTop(5)

	fmt.Println("Rand Top Test...")
	randData := data.New(10).Rand(10)
	top, err := randData.Top(5)
	fmt.Printf("\tData %v: %v\n\tTop %v: %v (err:%v)\n",
		10, *randData,
		5, top, err)

	fmt.Println("TestCase...")
	test := []int{9, 0, 1, 5, 2, 9, 2, 1, 2, 4}
	fmt.Println(test)
	(*top.TopData)(&test).ShowTop(4)
}

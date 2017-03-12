package main

import (
	"fmt"
	"math/cmplx"
	// "top"
	"topNew"
)

func add(x int, y int) int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

var (
	toBe          = false
	maxInt uint64 = 1<<64 - 1
	z             = cmplx.Sqrt(-5 + 12i)
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
	// fmt.Printf(f, toBe, toBe)
	// fmt.Printf(f, maxInt, maxInt)
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

	// arr := [...]int{2: 2, 3: 3}
	// fmt.Printf("arr: %T %v -> %T %v\n", arr, arr, arr[2:3], arr[2:3])
	// var slice = []int{1, 2, 3, 4, 5, 6}
	// ms := make([]int, 2, 5)
	// fmt.Println("arr", arr, len(arr), cap(arr))
	// fmt.Println("slice", slice, len(slice), cap(slice))
	// emptySlice := ms[1:1]
	// println("emptySlice", len(emptySlice), cap(emptySlice), emptySlice == nil)

	// //ms[1] = 1
	// // ms[4] = 4 // runtime error: index out of range
	// println("make slice", ms[0], ms[1], len(ms), cap(ms))

	// newS := slice[1:4]
	// fmt.Println("new slice", newS, len(newS), cap(newS))

	// ns := append(newS, 8, 8, 8)
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

	// 结构体的零值
	// type myStruct struct {
	// 	x, y int
	// }
	// var s myStruct
	// var p = new(myStruct)
	// fmt.Printf("myStruct s: %v\n", s)              // myStruct s: {0 0}
	// fmt.Printf("myStruct p: %v %v\n", p, p == nil) // myStruct p: &{0 0} false

	// slice的零值
	// var sliceInit top.TopData
	// fmt.Printf("sliceInit: %v %v\n", sliceInit, sliceInit == nil) // sliceInit: [] true

	fmt.Println("Seq Top Test...")
	var data top.TopData
	fmt.Printf("data: %v %v\n", data, data == nil) // data: [] true
	// data = top.TopData([]int{1, 1, 1, 1})
	// fmt.Printf("data: %v %v\n", data, data == nil)
	top.New(10 - 0).Seq().ShowTop(5)
	top.New(100000 - 0).Seq().ShowTop(100)
	top.New(100000 - 1).Seq().ShowTop(100)
	top.New(100000 - 2).Seq().ShowTop(100)

	fmt.Println("Rand Top Test...")
	randData := top.New(10).Rand(10)
	topData, err := randData.Top(5)
	fmt.Printf("\tData %v: %v\n\tTop %v: %v (err:%v)\n",
		10, *randData,
		5, topData, err)

	fmt.Println("TestCase...")
	test := []int{9, 0, 1, 5, 2, 9, 2, 1, 2, 4}
	fmt.Println(test)
	(*top.TopData)(&test).ShowTop(4)

	fmt.Print("\n更多的测试案例见 top_test.go\n\n")
}

// Test project main.go
package main

import (
	"fmt"
	"math"
)

// // 全局变量m
// var M = 100

func main() {
	//c1 := "H"
	//c2 := '你'
	//c3 := byte('H')
	//fmt.Printf("%T--%T--%T\n", c1, c2, c3)
	//fmt.Printf("%d--%d--%d\n", c1, c2, c3)

	//str := "ABC你好"
	//for i, v := range str {
	//	fmt.Println(i, v)
	//}
	//
	//for i, v := range []rune(str) {
	//	fmt.Println(i, v)
	//}

	//for i := 1; i <= 9; i++ {
	//	for j := i; j <= 9; j++ {
	//		fmt.Printf("%d*%d=%d \t", i, j, j*i)
	//	}
	//	fmt.Println()
	//}
	//fmt.Println()
	//for i := 9; i >= 1; i-- { //i 8
	//	for j := i; j <= 9; j++ { //j 8
	//		fmt.Printf("%d*%d=%d \t", i, j, j*i)
	//	}
	//	fmt.Println()
	//}
	//fmt.Println()
	//for i := 1; i <= 9; i++ {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d*%d=%d \t", j, i, j*i)
	//	}
	//	fmt.Println()
	//}
	//fmt.Println()
	//for i := 9; i >= 1; i-- {
	//	for j := 1; j <= i; j++ {
	//		fmt.Printf("%d*%d=%d \t", j, i, j*i)
	//	}
	//	fmt.Println()
	//}

	//fmt.Println(Reverse1("ABCD"))
	//fmt.Println(Reverse("你好AB"))

	fmt.Printf("%d \n", 100)
	fmt.Printf("%b \n", 100)
	fmt.Printf("%c \n", 44)
	fmt.Printf("%x \n", 17)
	fmt.Printf("%e \n", math.Pi)
}

func Reverse1(s string) string {
	b := []byte(s)
	for i := 0; i < len(b)/2; i++ {
		j := len(b) - i - 1
		b[i], b[j] = b[j], b[i]
	}
	return string(b)
}

func Reverse(s string) string {
	r := []rune(s)
	for i := 0; i < len(r)/2; i++ {
		j := len(r) - i - 1
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

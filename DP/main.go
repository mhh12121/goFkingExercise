package main

import "fmt"

type t struct {
	name string
	age  int
}

func main() {
	// prices := []int{7, 1, 5, 3, 6, 4}
	// fmt.Println("best Profit:", bestTime_stateMachine(prices))

	//---------slidewindow-----------------
	// t := []int{2, 3, 4, 2, 6, 2, 5, 1}
	t1 := []int{7, 6, 5, 4, 3, 2, 1}
	res := slideWindow(t1, 3)
	// fmt.Println("slide:", res)
	// t = t[:len(t)-1]
	// fmt.Println(t)
	//--------------perfectSquare-0-----------

	// fmt.Println("perfectnum:", PerfectSquare(10))
	// fmt.Println("perfectnum:", PerfectSquare_dp(10))

	//-------------lettercasePermutation--------------
	// s := make([]int, 0)
	// var tmp t
	// d := "a123b"
	// fmt.Println(letterCasePermutation(d))
	// do1(&tmp)
	// fmt.Println(tmp.age)
	// res := maxProfit(prices)
	fmt.Println(res)

}
func do1(stu *t) {

	stu.age = 1
}

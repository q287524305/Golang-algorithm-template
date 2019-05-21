package main

import "fmt"

func cmp(a, b []int) bool {
	if len(a) != len(b) {
		return len(a) > len(b)
	}
	for i := len(a) - 1; i >= 0; i-- {
		if a[i] != b[i] {
			return a[i] > b[i]
		}
	}
	return true
}

func bigNumSub(a, b []int) []int {
	res := []int{}
	t := 0
	for i := 0; i < len(a); i++ {
		t = a[i] - t
		if i < len(b) {
			t -= b[i]
		}
		res = append(res, (t+10)%10)
		if t < 0 {
			t = 1
		} else {
			t = 0
		}
	}
	for len(res) > 1 && res[len(res)-1] == 0 {
		res = res[:len(res)-1]
	}
	return res
}

func main() {
	a := "0"
	b := "0"
	arrA, arrB := []int{}, []int{}
	for i := len(a) - 1; i >= 0; i-- {
		arrA = append(arrA, int(a[i])-'0')
	}
	for i := len(b) - 1; i >= 0; i-- {
		arrB = append(arrB, int(b[i])-'0')
	}
	if cmp(arrA, arrB) {
		c := bigNumSub(arrA, arrB)
		for i := len(c) - 1; i >= 0; i-- {
			fmt.Print(c[i])
		}
	} else {
		fmt.Print("-")
		c := bigNumSub(arrB, arrA)
		for i := len(c) - 1; i >= 0; i-- {
			fmt.Print(c[i])
		}
	}

}

package main

import "fmt"

func main() {
	math := 80
	eng := 74
	history := 95
	fmt.Println("A님 평균 점수는", (math+eng+history)/3, "입니다.")

	math = 88
	eng = 92
	history = 53
	fmt.Println("B님 평균 점수는", (math+eng+history)/3, "입니다.")

	math = 78
	eng = 73
	history = 78
	fmt.Println("C님 평균 점수는", (math+eng+history)/3, "입니다.")

}

package main

import (
    "fmt"
    "time"
)

func Measure(repeat_cnt int,principal,goal float64){
	for i := 0; i <  repeat_cnt; i++ {
		var age_cnt int = 0
		var inner_principal float64 = principal
		for {
			inner_principal *= 1.05
			age_cnt++
			if inner_principal >= goal {
				fmt.Println(inner_principal)
				fmt.Println(age_cnt)
				break
			}
		}
	}
}

func main() {

	var (
		repeat_cnt int = 100
		principal float64 = 500000
		goal float64 = 45000000
	)

	start := time.Now()
	Measure(repeat_cnt,principal,goal)
	end := time.Now()
	fmt.Printf("end:%.17f[sec]\n",(end.Sub(start)).Seconds())
}


package main

import (
    "fmt"
    "time"
)

func main() {

	var pi4 float64 = 0
	start := time.Now()
	
	for i := 0; i < 100000000; i++ {
		pi4 += (1 / (float64(i) * 4 + 1) - 1 / (float64(i) * 4 + 3))
	}

	end := time.Now()
	fmt.Printf("end:%.17f[sec]\n",(end.Sub(start)).Seconds())
	fmt.Println(pi4 * 4)
}


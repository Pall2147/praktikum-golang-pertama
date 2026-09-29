package main

import("fmt"
	   "math"
	)
	   
func main() {
	var r float64
	fmt.Print("masukan r =")
	fmt.Scan(&r)
	L := math.Pi*r*r
	fmt.Printf("%.1f/n", L)
}
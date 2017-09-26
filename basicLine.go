package main

import "fmt"


func mean(xs[]float64)float64 {
	total:=0.0
	for _,v:=range xs {
		total +=v
	}
	return total/float64(len(xs))
}
func best_fit_slope(xs, ys, []int) float64{
	var m float64:= (((mean(xs)*mean(ys)) - mean(xs*ys)) / ((mean(xs)*mean(xs)) - mean(xs*xs)))
	return m
}

func main(){
	var xs[7] int
	var ys[7] int
	for i:=0; i < 7; i++ {
		ys[i] = i
		xs[i] = i
	}
	var m int := best_fit_slope(xs, ys)
	fmt.Println(m)
}

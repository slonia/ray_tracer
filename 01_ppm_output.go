package main

import "fmt"

func main() {
	nx := 200
	ny := 100
	fmt.Printf("P3\n%v %v\n255\n", nx, ny)
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			r := float32(i) / float32(nx)
			g := float32(j) / float32(ny)
			b := 0.2
			ir := int(255.99 * r)
			ig := int(255.99 * g)
			ib := int(255.99 * b)
			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
}

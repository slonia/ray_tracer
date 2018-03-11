package main

import (
	"fmt"
	"log"
)

func main() {
	nx := 600
	ny := 300
	maxColor := float32(255.99)
	fmt.Printf("P3\n%v %v\n255\n", nx, ny)
	lowerLeftCorner := vec3{-2.0, -1.0, -1.0}
	horizontal := vec3{4, 0.0, 0.0}
	vertical := vec3{0, 2.0, 0.0}
	origin := vec3{0.0, 0.0, 0.0}
	world := hitableList{[]*sphere{&sphere{center: vec3{0.0, 0.0, -1.0}, radius: 0.5}}, 1}
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float32(i) / float32(nx)
			v := float32(j) / float32(ny)
			r := ray{origin, lowerLeftCorner.plus(horizontal.multiplyBy(u)).plus(vertical.multiplyBy(v))}
			col := color(r, world)
			ir := int(maxColor * col.e0)
			ig := int(maxColor * col.e1)
			ib := int(maxColor * col.e2)
			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
}

func color(r ray, world hitable) vec3 {
	rec := &hitRecord{}

	if world.hit(r, 0.0, 100000000000.0, rec) {
		log.Printf("In main: %v\n", rec.normal)

		return vec3{rec.normal.x() + 1, rec.normal.y() + 1, rec.normal.z() + 1}.multiplyBy(0.5)
	} else {
		unit_direction := r.direction().unitVector()
		t := 0.5 * (unit_direction.y() + 1.0)
		return vec3{1.0, 1.0, 1.0}.multiplyBy(1.0 - t).plus(vec3{0.5, 0.7, 1.0}.multiplyBy(t))
	}
}

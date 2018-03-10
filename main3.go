package main

import (
	"fmt"
	"math"
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
	world := hitableList{[]hitRecord{sphere{vec3{0.0, 0.0, -1.0}, 0.5}, sphere{vec3{0.0, -100.5, -1.0}, 100}}, 2}
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

func color(r ray, world hitableList) vec3 {
	var rec hitRecord
	if world.hit(r, 0.0, 100000.0, rec) {
		return vec3{rec.normal.x() + 1, n.normal.y() + 1, n.normal.z() + 1}.multiplyBy(0.5)
	}
	unit_direction := r.direction().unitVector()
	t = 0.5 * (unit_direction.y() + 1.0)
	return vec3{1.0, 1.0, 1.0}.multiplyBy(1.0 - t).plus(vec3{0.5, 0.7, 1.0}.multiplyBy(t))
}

func hitSphere(center vec3, radius float32, r ray) float32 {
	oc := r.origin().minus(center)
	direction := r.direction()
	a := direction.dot(direction)
	b := 2.0 * oc.dot(direction)
	c := oc.dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	if discriminant < 0 {
		return -1.0
	} else {
		return (-b - float32(math.Sqrt(float64(discriminant)))) / (2.0 * a)
	}
}

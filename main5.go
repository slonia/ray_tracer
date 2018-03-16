package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	nx := 600
	ny := 300
	ns := 100
	maxColor := float32(255.99)
	cam := defaultCamera()
	fmt.Printf("P3\n%v %v\n255\n", nx, ny)
	world := hitableList{[]hitable{sphere{center: vec3{0.0, 0.0, -1.0}, radius: 0.5}, sphere{center: vec3{0.0, -100.5, -1.0}, radius: 100}}, 2}
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := vec3{0.0, 0.0, 0.0}
			for s := 0; s < ns; s++ {
				u := (float32(i) + rand.Float32()) / float32(nx)
				v := (float32(j) + rand.Float32()) / float32(ny)
				r := cam.getRay(u, v)
				col = col.plus(color(r, world))
			}
			col = col.divideBy(float32(ns))
			col = vec3{float32(math.Sqrt(float64(col.e0))), float32(math.Sqrt(float64(col.e1))), float32(math.Sqrt(float64(col.e2)))}
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
		target := rec.p.plus(rec.normal).plus(randomInUnitSphere())
		return color(ray{rec.p, target.minus(rec.p)}, world).multiplyBy(0.5)
	} else {
		unit_direction := r.direction().unitVector()
		t := 0.5 * (unit_direction.y() + 1.0)
		return vec3{1.0, 1.0, 1.0}.multiplyBy(1.0 - t).plus(vec3{0.5, 0.7, 1.0}.multiplyBy(t))
	}
}

func randomInUnitSphere() vec3 {
	var p vec3
	unitVec := vec3{1.0, 1.0, 1.0}
	for {
		p = vec3{rand.Float32(), rand.Float32(), rand.Float32()}.multiplyBy(2.0).minus(unitVec)
		if p.dot(p) < 1.0 {
			break
		}
	}
	return p
}

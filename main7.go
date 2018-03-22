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
	world := hitableList{[]hitable{
		sphere{center: vec3{0.0, 0.0, -1.0}, radius: 0.5, hitRecord: hitRecord{matPtr: lambertian{vec3{0.8, 0.3, 0.3}}}},
		sphere{center: vec3{0.0, -100.5, -1.0}, radius: 100, hitRecord: hitRecord{matPtr: lambertian{vec3{0.8, 0.8, 0.0}}}},
		sphere{center: vec3{1.0, 0.0, -1.0}, radius: 0.5, hitRecord: hitRecord{matPtr: metal{vec3{0.8, 0.6, 0.2}, 0.3}}},
		sphere{center: vec3{-1, 0.0, -1}, radius: 0.5, hitRecord: hitRecord{matPtr: dielectric{1.5}}},
		sphere{center: vec3{-1, 0.0, -1}, radius: -0.45, hitRecord: hitRecord{matPtr: dielectric{1.5}}}}, 4}
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			col := vec3{0.0, 0.0, 0.0}
			for s := 0; s < ns; s++ {
				u := (float32(i) + rand.Float32()) / float32(nx)
				v := (float32(j) + rand.Float32()) / float32(ny)
				r := cam.getRay(u, v)
				col = col.plus(color(r, world, 0))
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

func color(r ray, world hitable, depth int) vec3 {
	rec := &hitRecord{matPtr: lambertian{vec3{}}}

	if world.hit(r, 0.001, 100000000000.0, rec) {
		scattered := &ray{}
		attenuation := &vec3{}
		if depth < 50 && rec.matPtr.scatter(r, rec, attenuation, scattered) {
			return attenuation.multiply(color(*scattered, world, depth+1))
		} else {
			return vec3{0.0, 0.0, 0.0}
		}
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

func reflect(v vec3, n vec3) vec3 {
	return v.minus(n.multiplyBy(2.0 * v.dot(n)))
}

func refract(v vec3, n vec3, niOverNt float32, refracted *vec3) bool {
	uv := v.unitVector()
	dt := uv.dot(n)
	discriminant := 1 - niOverNt*niOverNt*(1.0-dt*dt)
	if discriminant > 0 {
		*refracted = v.minus(n.multiplyBy(dt)).multiplyBy(niOverNt).minus(n.multiplyBy(float32(math.Sqrt(float64(discriminant)))))
		return true
	}
	return false
}

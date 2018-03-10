package main

import "fmt"

func main() {
	nx := 600
	ny := 300
	max_color := float32(255.99)
	fmt.Printf("P3\n%v %v\n255\n", nx, ny)
	lower_left_corner := vec3{-2.0, -1.0, -1.0}
	horizontal := vec3{4, 0.0, 0.0}
	vertical := vec3{0, 2.0, 0.0}
	origin := vec3{0.0, 0.0, 0.0}
	for j := ny - 1; j >= 0; j-- {
		for i := 0; i < nx; i++ {
			u := float32(i) / float32(nx)
			v := float32(j) / float32(ny)
			r := ray{origin, lower_left_corner.plus(horizontal.multiply_by(u)).plus(vertical.multiply_by(v))}
			col := color(r)
			ir := int(max_color * col.e0)
			ig := int(max_color * col.e1)
			ib := int(max_color * col.e2)
			fmt.Printf("%v %v %v\n", ir, ig, ib)
		}
	}
}

func color(r ray) vec3 {
	if (hit_sphere(vec3{0.0, 0.0, -1.0}, 0.5, r)) {
		return vec3{1.0, 0.0, 0.0}
	}
	unit_direction := r.direction().unit_vector()
	t := 0.5 * (unit_direction.y() + 1.0)
	return vec3{1.0, 1.0, 1.0}.multiply_by(1.0 - t).plus(vec3{0.5, 0.7, 1.0}.multiply_by(t))
}

func hit_sphere(center vec3, radius float32, r ray) bool {
	oc := r.origin().minus(center)
	direction := r.direction()
	a := direction.dot(direction)
	b := 2.0 * oc.dot(direction)
	c := oc.dot(oc) - radius*radius
	discriminant := b*b - 4*a*c
	return discriminant > 0
}

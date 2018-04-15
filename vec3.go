package main

import "math"

type vec3 struct {
	e0, e1, e2 float64
}

func (vec vec3) x() float64 {
	return vec.e0
}

func (vec vec3) y() float64 {
	return vec.e1
}

func (vec vec3) z() float64 {
	return vec.e2
}

func (vec vec3) r() float64 {
	return vec.e0
}

func (vec vec3) g() float64 {
	return vec.e1
}

func (vec vec3) b() float64 {
	return vec.e2
}

func (vec vec3) squaredLength() float64 {
	return vec.e0*vec.e0 + vec.e1*vec.e1 + vec.e2*vec.e2
}

func (vec vec3) length() float64 {
	return math.Sqrt(vec.squaredLength())
}

func (v1 vec3) plus(v2 vec3) vec3 {
	return vec3{v1.e0 + v2.e0, v1.e1 + v2.e1, v1.e2 + v2.e2}
}

func (v1 vec3) minus(v2 vec3) vec3 {
	return vec3{v1.e0 - v2.e0, v1.e1 - v2.e1, v1.e2 - v2.e2}
}

func (v1 vec3) multiply(v2 vec3) vec3 {
	return vec3{v1.e0 * v2.e0, v1.e1 * v2.e1, v1.e2 * v2.e2}
}

func (v1 vec3) divide(v2 vec3) vec3 {
	return vec3{v1.e0 / v2.e0, v1.e1 / v2.e1, v1.e2 / v2.e2}
}

func (v1 vec3) multiplyBy(f float64) vec3 {
	return vec3{v1.e0 * f, v1.e1 * f, v1.e2 * f}
}

func (v1 vec3) divideBy(f float64) vec3 {
	return vec3{v1.e0 / f, v1.e1 / f, v1.e2 / f}
}

func (v1 vec3) dot(v2 vec3) float64 {
	return v1.e0*v2.e0 + v1.e1*v2.e1 + v1.e2*v2.e2
}

func (v1 vec3) cross(v2 vec3) vec3 {
	return vec3{v1.e1*v2.e2 - v1.e2*v2.e1,
		-(v1.e0*v2.e2 - v1.e2*v2.e0),
		v1.e0*v2.e1 - v1.e1*v2.e0}
}

func (v1 vec3) unitVector() vec3 {
	return v1.divideBy(v1.length())
}

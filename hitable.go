package main

import "math"

type hitRecord struct {
	t         float32
	p, normal vec3
}

type hitable interface {
	hit(r ray, tMin float32, tMax float32, rec hitRecordInterface) bool
}

type sphere struct {
	hitRecord
	center vec3
	radius float32
}

type hitRecordInterface interface {
	setT(t float32)
	setP(vec vec3)
	setNormal(vec vec3)
	getT() float32
	getP() vec3
	getNormal() vec3
}

func (s *sphere) setT(t float32) {
	s.t = t
}

func (s *sphere) setP(vec vec3) {
	s.p = vec
}

func (s *sphere) setNormal(vec vec3) {
	s.normal = vec
}

func (s sphere) getT() float32 {
	return s.t
}

func (s sphere) getP() vec3 {
	return s.p
}

func (s sphere) getNormal() vec3 {
	return s.normal
}

func (s *hitRecord) setT(t float32) {
	s.t = t
}

func (s *hitRecord) setP(vec vec3) {
	s.p = vec
}

func (s *hitRecord) setNormal(vec vec3) {
	s.normal = vec
}

func (s hitRecord) getT() float32 {
	return s.t
}

func (s hitRecord) getP() vec3 {
	return s.p
}

func (s hitRecord) getNormal() vec3 {
	return s.normal
}

func (s sphere) hit(r ray, tMin float32, tMax float32, rec hitRecordInterface) bool {
	oc := r.origin().minus(s.center)
	direction := r.direction()
	a := direction.dot(direction)
	b := oc.dot(direction)
	c := oc.dot(oc) - s.radius*s.radius
	discriminant := b*b - a*c
	if discriminant > 0 {
		temp := (-b - float32(math.Sqrt(float64(discriminant)))) / a
		if temp < tMax && temp > tMin {
			rec.setT(temp)
			rec.setP(r.pointAtParameter(temp))
			rec.setNormal(rec.getP().minus(s.center).divideBy(s.radius))
			return true
		}
		temp = (-b + float32(math.Sqrt(float64(discriminant)))) / a
		if temp < tMax && temp > tMin {
			rec.setT(temp)
			rec.setP(r.pointAtParameter(temp))
			rec.setNormal(rec.getP().minus(s.center).divideBy(s.radius))
			return true
		}
	}
	return false
}

type hitableList struct {
	list     []hitable
	listSize int
}

func (hl hitableList) hit(r ray, tMin float32, tMax float32, rec hitRecordInterface) bool {
	tempRec := &hitRecord{0, vec3{0.0, 0.0, 0.0}, vec3{0.0, 0.0, 0.0}}
	hitAnything := false
	closestSoFar := tMax
	for i := 0; i < hl.listSize; i++ {
		if hl.list[i].hit(r, tMin, closestSoFar, tempRec) {
			hitAnything = true
			closestSoFar = tempRec.getT()
			rec.setT(tempRec.getT())
			rec.setP(tempRec.getP())
			rec.setNormal(tempRec.getNormal())
		}
	}
	return hitAnything
}

package main

import (
	"log"
	"math"
)

type hitRecord struct {
	t      float32
	p      vec3
	normal vec3
}

type hitable interface {
	hit(r ray, tMin float32, tMax float32, rec *hitRecord) bool
}

type sphere struct {
	hitRecord
	center vec3
	radius float32
}

func (h hitRecord) hit(r ray, tMin float32, tMax float32, rec *hitRecord) bool {
	return false
}

func (s sphere) hit(r ray, tMin float32, tMax float32, rec *hitRecord) bool {
	oc := r.origin().minus(s.center)
	direction := r.direction()
	a := direction.dot(direction)
	b := 2.0 * oc.dot(direction)
	c := oc.dot(oc) - s.radius*s.radius
	discriminant := b*b - 4*a*c
	if discriminant > 0 {
		temp := (-b - float32(math.Sqrt(float64(b*b-a*c)))) / a
		if temp < tMax && temp > tMax {
			rec.t = temp
			rec.p = r.pointAtParameter(rec.t)
			rec.normal = rec.p.minus(s.center).divideBy(s.radius)
			return true
		}
		temp = (-b + float32(math.Sqrt(float64(b*b-a*c)))) / a
		if temp < tMax && temp > tMin {
			rec.t = temp
			rec.p = r.pointAtParameter(rec.t)
			rec.normal = rec.p.minus(s.center).divideBy(s.radius)
			return true
		}
	}
	return false
}

type hitableList struct {
	hitable  []*sphere
	listSize int
}

func (list hitableList) hit(r ray, tMin float32, tMax float32, rec *hitRecord) bool {
	// log.Printf("In hit1: %p\n", rec)
	var tempRec hitRecord
	hitAnything := false
	closestSoFar := tMax
	for i := 0; i < list.listSize; i++ {
		if list.hitable[i].hit(r, tMin, closestSoFar, &tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			rec = &tempRec
			log.Printf("In hit: %p\n", rec)
			log.Printf("In hit: %v\n", *rec)
		}
	}
	return hitAnything
}

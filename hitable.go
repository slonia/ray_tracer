package main

import (
	"math"
)

type hitRecord struct {
	t         float32
	p, normal vec3
	matPtr    material
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
	getHR() hitRecord
	setHR(hitRecord)
}

func (s sphere) getHR() hitRecord {
	return s.hitRecord
}

func (hr hitRecord) getHR() hitRecord {
	return hr
}

func (s *sphere) setHR(h hitRecord) {
	s.hitRecord.setHR(h)
}

func (hr *hitRecord) setHR(h hitRecord) {
	*hr = h
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
			hr := rec.getHR()
			hr.t = temp
			hr.p = r.pointAtParameter(temp)
			hr.normal = hr.p.minus(s.center).divideBy(s.radius)
			hr.matPtr = s.matPtr
			rec.setHR(hr)
			return true
		}
		temp = (-b + float32(math.Sqrt(float64(discriminant)))) / a
		if temp < tMax && temp > tMin {
			hr := rec.getHR()
			hr.t = temp
			hr.p = r.pointAtParameter(temp)
			hr.normal = hr.p.minus(s.center).divideBy(s.radius)
			hr.matPtr = s.matPtr
			rec.setHR(hr)
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
	tempRec := &hitRecord{t: 0, p: vec3{0.0, 0.0, 0.0}, normal: vec3{0.0, 0.0, 0.0}, matPtr: lambertian{vec3{0.0, 0.0, 0.0}}}
	hitAnything := false
	closestSoFar := tMax
	for _, item := range hl.list {
		if item.hit(r, tMin, closestSoFar, tempRec) {
			hitAnything = true
			hr := tempRec.getHR()
			closestSoFar = hr.t
			rec.setHR(hr)
		}
	}
	return hitAnything
}

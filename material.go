package main

import (
	"math"
	"math/rand"
)

type material interface {
	scatter(rIn ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool
}

type lambertian struct {
	albedo vec3
}

func (l lambertian) scatter(rIn ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool {
	target := rec.p.plus(rec.normal).plus(randomInUnitSphere())
	*scattered = ray{rec.p, target.minus(rec.p)}
	*attenuation = l.albedo
	return true
}

type metal struct {
	albedo vec3
	fuzz   float32
}

func (m metal) scatter(rIn ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool {
	reflected := reflect(rIn.direction().unitVector(), rec.normal)
	fuzz := m.fuzz
	if fuzz > 1.0 {
		fuzz = 1.0
	}
	*scattered = ray{rec.p, reflected.plus(randomInUnitSphere().multiplyBy(fuzz))}
	*attenuation = m.albedo
	return scattered.direction().dot(rec.normal) > 0
}

type dielectric struct {
	refIdx float32
}

func (d dielectric) scatter(rIn ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool {
	outwardNormal := vec3{0.0, 0.0, 0.0}
	refracted := vec3{0.0, 0.0, 0.0}
	var niOverNt, reflectProb, cosine float32
	reflected := reflect(rIn.direction(), rec.normal)
	*attenuation = vec3{1.0, 1.0, 1.0}
	if rIn.direction().dot(rec.normal) > 0 {
		outwardNormal = vec3{-rec.normal.e0, -rec.normal.e1, -rec.normal.e2}
		niOverNt = d.refIdx
		cosine = niOverNt * rIn.direction().dot(rec.normal) / rIn.direction().length()
	} else {
		outwardNormal = rec.normal
		niOverNt = 1.0 / d.refIdx
		cosine = -rIn.direction().dot(rec.normal) / rIn.direction().length()
	}
	if refract(rIn.direction(), outwardNormal, niOverNt, &refracted) {
		reflectProb = schlick(cosine, d.refIdx)
	} else {
		reflectProb = 1.0
	}

	if rand.Float32() < reflectProb {
		*scattered = ray{rec.p, reflected}
	} else {
		*scattered = ray{rec.p, refracted}
	}

	return true
}

func schlick(cosine float32, refIdx float32) float32 {
	r0 := (1 - refIdx) / (1 + refIdx)
	r0 *= r0
	return r0 + (1-r0)*float32(math.Pow(float64(1-cosine), 5))
}

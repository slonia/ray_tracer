package main

type material interface {
	scatter(rIn ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool
}

type lambertian struct {
	albedo vec3
}

func (l lambertian) scatter(rIn ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool {
	target := rec.p.plus(rec.normal).plus(randomInUnitSphere())
	scattered = &ray{rec.p, target.minus(rec.p)}
	attenuation = &l.albedo
	return true
}

type metal struct {
	albedo vec3
}

func (m metal) scatter(rIn ray, rec *hitRecord, attenuation *vec3, scattered *ray) bool {
	reflected := reflect(rIn.direction().unitVector(), rec.normal)
	scattered = &ray{rec.p, reflected}
	attenuation = &m.albedo
	return scattered.direction().dot(rec.normal) > 0
}

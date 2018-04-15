package main

import (
	"math"
	"math/rand"
)

type camera struct {
	origin, lowerLeftCorner, horizontal, vertical, u, v, w vec3
	lensRadius                                             float64
}

func (c camera) getRay(u float64, v float64) ray {
	rd := randomInUnitDisk().multiplyBy(c.lensRadius)
	offset := c.u.multiplyBy(rd.x()).plus(c.v.multiplyBy(rd.y()))
	return ray{c.origin.plus(offset), c.lowerLeftCorner.plus(c.horizontal.multiplyBy(u)).plus(c.vertical.multiplyBy(v).minus(c.origin)).minus(offset)}
}

func defaultCamera(lookfrom vec3, lookat vec3, vup vec3, vfov float64, aspect float64, aperture float64, focusDist float64) *camera {
	lensRadius := aperture / 2.0
	theta := vfov * math.Pi / 180.0
	halfHeight := math.Tan(theta / 2.0)
	halfWidth := aspect * halfHeight
	w := lookfrom.minus(lookat).unitVector()
	u := vup.cross(w).unitVector()
	v := w.cross(u)
	llCorner := lookfrom.minus(u.multiplyBy(halfWidth * focusDist)).minus(v.multiplyBy(halfHeight * focusDist)).minus(w.multiplyBy(focusDist))
	return &camera{
		origin:          lookfrom,
		lowerLeftCorner: llCorner,
		horizontal:      u.multiplyBy(2.0 * halfWidth * focusDist),
		vertical:        v.multiplyBy(2.0 * halfHeight * focusDist),
		u:               u,
		v:               v,
		w:               w,
		lensRadius:      lensRadius}
}

func randomInUnitDisk() vec3 {
	var p vec3
	for {
		p = vec3{rand.Float64(), rand.Float64(), 0}.multiplyBy(2.0).minus(vec3{1.0, 1.0, 0.0})
		if p.dot(p) < 1.0 {
			break
		}
	}
	return p
}

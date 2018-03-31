package main

import "math"

type camera struct {
	origin, lowerLeftCorner, horizontal, vertical vec3
}

func (c camera) getRay(u float32, v float32) ray {
	return ray{c.origin, c.lowerLeftCorner.plus(c.horizontal.multiplyBy(u)).plus(c.vertical.multiplyBy(v).minus(c.origin))}
}

func defaultCamera(lookfrom vec3, lookat vec3, vup vec3, vfov float32, aspect float32) *camera {
	theta := vfov * math.Pi / 180.0
	halfHeight := float32(math.Tan(float64(theta / 2.0)))
	halfWidth := aspect * halfHeight
	w := lookfrom.minus(lookat).unitVector()
	u := vup.cross(w).unitVector()
	v := w.cross(u)
	llCorner := lookfrom.minus(u.multiplyBy(halfWidth)).minus(v.multiplyBy(halfHeight)).minus(w)
	return &camera{
		origin:          lookfrom,
		lowerLeftCorner: llCorner,
		horizontal:      u.multiplyBy(2.0 * halfWidth),
		vertical:        v.multiplyBy(2.0 * halfHeight)}
}

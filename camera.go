package main

type camera struct {
	origin, lowerLeftCorner, horizontal, vertical vec3
}

func (c camera) getRay(u float32, v float32) ray {
	return ray{c.origin, c.lowerLeftCorner.plus(c.horizontal.multiplyBy(u)).plus(c.vertical.multiplyBy(v))}
}

func defaultCamera() *camera {
	return &camera{vec3{0.0, 0.0, 0.0}, vec3{-2.0, -1.0, -1.0}, vec3{4.0, 0.0, 0.0}, vec3{0.0, 2.0, 0.0}}
}

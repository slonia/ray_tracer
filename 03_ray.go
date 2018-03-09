package main

type ray struct {
	A, B vec3
}

func (r ray) origin() vec3 {
	return r.A
}

func (r ray) direction() vec3 {
	return r.B
}

func (r ray) point_at_parameter(t float32) vec3 {
	return r.A.plus(r.B.multiply_by(t))
}

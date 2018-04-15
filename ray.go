package main

type ray struct {
	a, b vec3
}

func (r ray) origin() vec3 {
	return r.a
}

func (r ray) direction() vec3 {
	return r.b
}

func (r ray) pointAtParameter(t float64) vec3 {
	return r.a.plus(r.b.multiplyBy(t))
}

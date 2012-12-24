package main

import (
	"rk4"
	. "vector"
)

type body struct {
	pos, vel, forces, input Vector
	mass             float64
}

type phobject interface {
	force_func(x, p Vector, t float64) Vector
	update_vecs(new_x, new_p Vector)
	getx() Vector
	getp() Vector
}

func (b *body) force_func(x, p Vector, t float64) Vector {
	return p.Scale(-0.05).Add(b.forces.Scale(1e-3))
}

func (b *body) update_vecs(new_x, new_p Vector) {
	b.pos = new_x
	b.vel = new_p
}

func (b *body) getx() Vector {
	return b.pos
}

func (b *body) getp() Vector {
	return b.vel
}

func curry_force(b phobject) func(Vector, Vector, float64) Vector {
	return func(x, y Vector, t float64) Vector {
		return b.force_func(x, y, t)
	}
}

func time_step(b phobject, t, dt float64) {
	b.update_vecs(rk4.Integrate(b.getx(), b.getp(), t, dt, curry_force(b)))
}

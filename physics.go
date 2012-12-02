package main

type body struct {
	pos, vel, accel vector
	mass float64
}

func (b *body) time_step(dt float64) {
	b.vel = b.vel.add(b.accel).scale(dt)
	b.pos = b.pos.add(b.vel.scale(dt*0.001))
}

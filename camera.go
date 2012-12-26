package main

import (
	. "vector"
)

type camera struct {
	*body
	top, front, target Vector
}

func (c *camera) init(v Vector) {
	c.body = new(body)
	c.pos = v

	c.top.X = c.pos.X
	c.top.Y = c.pos.Y
	c.top.Z = c.pos.Z + 1.0

	c.front.X = c.pos.X + 1.0
	c.front.Y = c.pos.Y
	c.front.Z = c.pos.Z

	c.mass = 1
	//c.speed = 1
}

func (c *camera) move(v Vector) {
	c.pos.X = v.X
	c.pos.Y = v.Y
	c.pos.Z = v.Z

	/*
		c.top.X = v.X
		c.top.Y = v.Y
		c.top.Z = v.Z  + 1.0
	*/
	//c.front.X = v.X
	//c.front.Y = v.Y + 1.0
	//c.front.Z = v.Z
}

//As in point the camera at something.
//I'm assuming the camera is at some negative Y, pointing toward positive Y with positive X on the right and positive Z up.
func (c *camera) point_at(v Vector) {
  c.front = v.Sub(c.pos).Normalize().Add(c.pos)
}

func (c *camera) update(dt float64) {
	c.forces = c.input.Normalize()
	time_step(c, 0, dt)
	c.target = c.pos.Add(Vector{0, 1, 0})
	c.point_at(c.target)
}

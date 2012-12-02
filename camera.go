package main

type camera struct {
	body
	top, front vector
}

func (c *camera) init(v vector) {
	c.pos.x = v.x
	c.pos.y = v.y
	c.pos.z = v.z

	c.top.x = c.pos.x
	c.top.y = c.pos.y
	c.top.z = c.pos.z + 1.0

	c.front.x = c.pos.x + 1.0
	c.front.y = c.pos.y
	c.front.z = c.pos.z

	c.mass = 1
	//c.speed = 1
}

func (c *camera) move(v vector) {
	c.pos.x = v.x
	c.pos.y = v.y
	c.pos.z = v.z

	/*
	c.top.x = v.x
	c.top.y = v.y
	c.top.z = v.z  + 1.0
	*/
	//c.front.x = v.x
	//c.front.y = v.y + 1.0
	//c.front.z = v.z
}

//As in point the camera at something.
//I'm assuming the camera is at some negative y, pointing toward positive y with positive x on the right and positive z up.
func (c *camera) point_at(v vector) {
	c.front.x = vector{x:v.x - c.pos.x, y:v.y-c.pos.y, z:v.z-c.pos.z}.normalize().x// + c.x//Relative vector
	c.front.y = vector{x:v.x - c.pos.x, y:v.y-c.pos.y, z:v.z-c.pos.z}.normalize().y// + c.y
	c.front.z = vector{x:v.x - c.pos.x, y:v.y-c.pos.y, z:v.z-c.pos.z}.normalize().z// + c.z
}

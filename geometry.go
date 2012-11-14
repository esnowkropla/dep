package main

import (
	"github.com/banthar/gl"
	"math"
)

type vector struct {
	x, y, z float64
}

type sprite struct {
	vector
	height, width float32
	tex gl.Texture
}

func (u vector) dot(v vector) vector  {
	return vector{x:u.x*v.x, y:u.y*v.y, z:u.z*v.z}
}

func (u vector) cross(v vector) vector {
	return vector{x:u.y*v.z - u.z*v.y, y:u.z*v.x - u.x*v.z, z:u.x*v.y - u.y*v.x}
}

func (v vector) normalize() vector {
	mag := math.Sqrt(v.x*v.x + v.y*v.y + v.z*v.z)
	return vector{x:v.x/mag, y:v.y/mag, z:v.z/mag}
}

func (u vector) add(v vector) vector {
	return vector{x:u.x+v.x, y:u.y+v.y, z:u.z+v.z}
}

func (u vector) sub(v vector) vector {
	return vector{x:u.x-v.x, y:u.y-v.y, z:u.z-v.z}
}

func make_vector(a, b, c float64) vector {
	return vector{x:a, y:b, z:c}
}

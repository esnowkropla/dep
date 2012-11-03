package main

import (
	"github.com/banthar/gl"
)

type vertex struct {
	x, y, z float64
}

type vector struct {
	pos vertex
}

type sprite struct {
	vertex
	height, width float32
	tex gl.Texture
}

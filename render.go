package main

import (
	"github.com/banthar/gl"
	"github.com/banthar/glu"
	"container/list"
)

type sprite_list struct {
	list.List
}

func resize_window(x, y int) {
	gl.Viewport(0, 0, x, y)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	glu.Perspective(45.0, float64(x)/float64(y), 1.0, 1000.0)
}

func (s sprite) render () {
  gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)

	gl.PushMatrix()
	gl.Translatef(float32(s.x), float32(s.y), float32(s.z))

	s.tex.Bind(gl.TEXTURE_2D)
	//gl.Color3f(0.3, 0.3, 0.3)
	gl.Begin(gl.QUADS)
		gl.TexCoord2f(0.0, 1.0); gl.Vertex3f(-s.width/2.0, -s.height/2.0, 0.0)
		gl.TexCoord2f(1.0, 1.0); gl.Vertex3f(s.width/2.0, -s.height/2.0, 0.0)
		gl.TexCoord2f(1.0, 0.0); gl.Vertex3f(s.width/2.0, s.height/2.0, 0.0)
		gl.TexCoord2f(0.0, 0.0); gl.Vertex3f(-s.width/2.0, s.height/2.0, 0.0)
	gl.End()
	s.tex.Unbind(gl.TEXTURE_2D)
	gl.PopMatrix()
}


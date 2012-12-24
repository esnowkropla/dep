package main

import (
	"container/list"
	"github.com/banthar/gl"
	"github.com/banthar/glu"
	. "vector"
)

type sprite struct {
	Vector
	height, width float32
	tex           gl.Texture
}

func size_window(x, y int, fov, aspect, min_cull, max_cull float64) {
	gl.Viewport(0, 0, x, y)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	glu.Perspective(fov, aspect, min_cull, max_cull)
	gl.Enable(gl.DEPTH_TEST)
}

func resize_window(x, y int) {
	size_window(x, y, 45.0, float64(x)/float64(y), 1.0, 1000.0)
}

func general_render(s *list.List, c *camera) {
	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.PushMatrix()
	//gl.Translatef(float32(c.X), float32(c.Y), float32(c.Z))
	glu.LookAt(c.pos.X, c.pos.Y, c.pos.Z, c.front.X, c.front.Y, c.front.Z, c.top.X, c.top.Y, c.top.Z)

	for e := s.Front(); e != nil; e = e.Next() {
		gl.PushMatrix()
		v := e.Value.(*sprite)
		v.render()
		gl.PopMatrix()
	}
	gl.PopMatrix()
}

func (s sprite) render() {
	gl.Translatef(float32(s.X), float32(s.Y), float32(s.Z))
	gl.Rotatef(90.0, 1.0, 0.0, 0.0)

	s.tex.Bind(gl.TEXTURE_2D)
	gl.Begin(gl.QUADS)
	gl.TexCoord2f(0.0, 1.0)
	gl.Vertex3f(-s.width/2.0, -s.height/2.0, 0.0)
	gl.TexCoord2f(1.0, 1.0)
	gl.Vertex3f(s.width/2.0, -s.height/2.0, 0.0)
	gl.TexCoord2f(1.0, 0.0)
	gl.Vertex3f(s.width/2.0, s.height/2.0, 0.0)
	gl.TexCoord2f(0.0, 0.0)
	gl.Vertex3f(-s.width/2.0, s.height/2.0, 0.0)
	gl.End()
	s.tex.Unbind(gl.TEXTURE_2D)
}

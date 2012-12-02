package main

import (
	"github.com/banthar/gl"
	"github.com/banthar/glu"
	"container/list"
)

type sprite_list struct {
	list.List
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

func general_render(s *sprite_list, c *camera) {

	gl.Clear(gl.COLOR_BUFFER_BIT | gl.DEPTH_BUFFER_BIT)
	gl.PushMatrix()
	//gl.Translatef(float32(c.x), float32(c.y), float32(c.z))
	glu.LookAt(c.pos.x, c.pos.y, c.pos.z, c.front.x, c.front.y, c.front.z, c.top.x, c.top.y, c.top.z)

	for e := s.Front(); e != nil; e = e.Next() {
		gl.PushMatrix()
		v := e.Value.(*sprite)
		v.render()
		gl.PopMatrix()
	}
	gl.PopMatrix()
}

func (s sprite) render () {
	gl.Translatef(float32(s.x), float32(s.y), float32(s.z))
	gl.Rotatef(90.0, 1.0, 0.0, 0.0)

	s.tex.Bind(gl.TEXTURE_2D)
	gl.Begin(gl.QUADS)
		gl.TexCoord2f(0.0, 1.0); gl.Vertex3f(-s.width/2.0, -s.height/2.0, 0.0)
		gl.TexCoord2f(1.0, 1.0); gl.Vertex3f(s.width/2.0, -s.height/2.0, 0.0)
		gl.TexCoord2f(1.0, 0.0); gl.Vertex3f(s.width/2.0, s.height/2.0, 0.0)
		gl.TexCoord2f(0.0, 0.0); gl.Vertex3f(-s.width/2.0, s.height/2.0, 0.0)
	gl.End()
	s.tex.Unbind(gl.TEXTURE_2D)
}


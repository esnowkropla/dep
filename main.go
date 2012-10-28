package main

import (
	"github.com/banthar/gl"
	"github.com/jteeuwen/glfw"
	//"time"
	//"fmt"
)

func main() {
	type glob struct {
		vertex_buffer, element_buffer gl.VertexArray
	}

	w := 640
	h := 480

	//aspect_ratio := float64(w)/float64(h)

	glfw.Init()
	glfw.OpenWindow(w, h, 8, 8, 8, 8, 8, 0, glfw.Windowed)
	glfw.SetWindowTitle("Shader Test")

	running := true

	for running {
		//update()
		//render()

		glfw.SwapBuffers()
		running = glfw.Key(glfw.KeyEsc) == 0 && glfw.WindowParam(glfw.Opened) != 0
	}
}

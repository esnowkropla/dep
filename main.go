package main

import (
	"math"
	"github.com/banthar/gl"
	"github.com/banthar/glu"
	"github.com/jteeuwen/glfw"
	//"time"
	"fmt"
	//"image"
	//"image/gif"
	"image/png"
	//"image/color"
	"os"
	//"github.com/vova616/GarageEngine/Engine"
	//"math/rand"
	//"errors"
)

func die (e error) {
	fmt.Println(e)
	os.Exit(1)
}

func texture(s string) gl.Texture {
	f, err := os.Open(s)
	if err != nil {die(err)}

	img, err := png.Decode(f)
	if err != nil {die(err)	}

	w := img.Bounds().Dx()
	h := img.Bounds().Dy()

	gl.Enable(gl.TEXTURE_2D)
	tex := gl.GenTexture()
	tex.Bind(gl.TEXTURE_2D)
	//gl.PixelStorei(gl.UNPACK_ALIGNMENT, 1)

	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_S, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_WRAP_T, gl.REPEAT)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MAG_FILTER, gl.LINEAR)
	gl.TexParameteri(gl.TEXTURE_2D, gl.TEXTURE_MIN_FILTER, gl.LINEAR)

	//gl.TexEnvf(gl.TEXTURE_ENV, gl.TEXTURE_ENV_MODE, gl.MODULATE)

	internal_format, img_type, format, target, err := ColorModelToGLTypes(img.ColorModel())
	if err != nil {die(err)}

	data, err := ImageData(img)
	if err != nil {die(err)}

	gl.TexImage2D(target, 0, internal_format, w, h, 0, img_type, format, data)
	tex.Unbind(gl.TEXTURE_2D)
	return tex
}

func main() {
	type glob struct {
		vertex_buffer, element_buffer gl.VertexArray
	}

	w := 500
	h := 500

	//aspect_ratio := float64(w)/float64(h)

	glfw.Init()
	glfw.OpenWindow(w, h, 8, 8, 8, 8, 8, 0, glfw.Windowed)
	glfw.SetWindowTitle("Shader Test")
	glfw.SetWindowSizeCallback(resize_window)

  gl.ClearColor(0.3, 0.3, 0.3, 1.0)
	gl.ShadeModel(gl.SMOOTH)
	gl.Viewport(0, 0, w, h)
	gl.MatrixMode(gl.PROJECTION)
	gl.LoadIdentity()
	glu.Perspective(45.0, float64(w)/float64(h), 1.0, 1000000.0)
	gl.Enable(gl.DEPTH_TEST)
	gl.PolygonMode(gl.FRONT_AND_BACK, gl.FILL)

	test_sprite := sprite{height:1.0, width:1.0}
	test_sprite.tex = texture("./assets/texture.png")
	test_sprite.z = -3
	var t float64 = 0

	new_guy := sprite{height:0.5, width:0.5}
	new_guy.tex = texture("./assets/red.png")
	new_guy.z = -5

	renderees := new(sprite_list)
	renderees.PushBack(new_guy)
	renderees.PushBack(test_sprite)

	running := true
	for running {
		//update()
		t += 0.001
		test_sprite.x = math.Cos(t)
		test_sprite.y = math.Sin(t)
		test_sprite.render()

		glfw.SwapBuffers()
		running = glfw.Key(glfw.KeyEsc) == 0 && glfw.WindowParam(glfw.Opened) != 0
	}
}

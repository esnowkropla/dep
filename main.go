package main

import (
	"math"
	"github.com/banthar/gl"
	//"github.com/banthar/glu"
	"github.com/jteeuwen/glfw"
	//"time"
	"fmt"
	//"image"
//_	"image/gif"
//_	"image/png"
//_	"image/jpeg"
	//"image/color"
	"os"
	//"github.com/vova616/GarageEngine/Engine"
	//"math/rand"
	//"errors"
	//"flag"
	//"runtime/pprof"
	//"log"
)

var my_camera = new(camera)

func die (e error) {
	fmt.Println(e)
	os.Exit(1)
}

func main() {
	w := 500
	h := 500
	fov := 45.0
	clip_min := 0.1
	clip_max := 1000.0

	aspect_ratio := float64(w)/float64(h)

	glfw.Init()
	glfw.OpenWindow(w, h, 8, 8, 8, 8, 8, 0, glfw.Windowed)

	glfw.SetWindowTitle("Shader Test")
	glfw.SetWindowSizeCallback(resize_window)
	glfw.SetKeyCallback(key_callback)

  gl.ClearColor(0.3, 0.3, 0.3, 1.0)
	gl.ShadeModel(gl.SMOOTH)

	size_window(w, h, fov, aspect_ratio, clip_min, clip_max)

	//Testing sprites
	test_sprite := sprite{height:1.0, width:1.0}
	test_sprite.tex = texture("./assets/hedge.gif")
	test_sprite.y = 3
	var t float64 = 0

	new_guy := sprite{height:0.5, width:0.5}
	new_guy.tex = texture("./assets/red2.png")
	new_guy.y = 5

	renderees := new(sprite_list)
	renderees.PushBack(&new_guy)
	renderees.PushBack(&test_sprite)

	my_camera.init(vector{x:0.0, y:-1.0, z:0.0})
	my_camera.front = vector{x:0.0, y:1.0, z:0.0}
	my_camera.top = vector{x:0.0, y:0.0, z:1.0}
	//End testing sprites

	running := true
	for running {
		//update()
		t += 0.001
		new_guy.x = math.Cos(t)
		new_guy.y = math.Sin(t) + 3.0
		//new_guy.z = math.Sin(t)
		//my_camera.move(vector{x:new_guy.x, y:0.0, z:0.0})
		//my_camera.x = new_guy.x
		//my_camera.point_at(vector{x:new_guy.x, y:new_guy.y, z:new_guy.z})

		general_render(renderees, my_camera)

		glfw.SwapBuffers()
		running = glfw.Key(glfw.KeyEsc) == 0 && glfw.WindowParam(glfw.Opened) != 0
	}
}

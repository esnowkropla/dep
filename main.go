package main

import (
	"github.com/banthar/gl"
	"math"
	//"github.com/banthar/glu"
	"fmt"
	"github.com/jteeuwen/glfw"
	"time"
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
	"container/list"
	. "vector"
)

var my_camera = new(camera)

func die(e error) {
	fmt.Println(e)
	os.Exit(1)
}

func update(l *list.List, dt float64) {
	for e := l.Front(); e != nil; e = e.Next() {
		switch e.Value.(type) {
		case *camera: //Could implement this with an interface implementing e.g. IsBody
			e.Value.(*camera).update(dt)
		case *body:
			time_step(*e.Value.(*phobject), 0, dt)
		}
	}
}

func main() {
	w := 500
	h := 500
	fov := 45.0
	clip_min := 0.1
	clip_max := 1000.0

	aspect_ratio := float64(w) / float64(h)

	glfw.Init()
	glfw.OpenWindow(w, h, 8, 8, 8, 8, 8, 0, glfw.Windowed)

	glfw.SetWindowTitle("Shader Test")
	glfw.SetWindowSizeCallback(resize_window)
	glfw.SetKeyCallback(key_callback)

	gl.ClearColor(0.3, 0.3, 0.3, 1.0)
	gl.ShadeModel(gl.SMOOTH)

	size_window(w, h, fov, aspect_ratio, clip_min, clip_max)

	//Testing sprites
	test_sprite := sprite{height: 1.0, width: 1.0}
	test_sprite.tex = texture("./assets/hedge.gif")
	test_sprite.Y = 3

	new_guy := sprite{height: 0.5, width: 0.5}
	new_guy.tex = texture("./assets/red2.png")
	new_guy.Y = 5

	renderees := new(list.List)
	renderees.PushBack(&new_guy)
	renderees.PushBack(&test_sprite)

	my_camera.init(Vector{0.0, -1.0, 0.0})
	my_camera.front = Vector{0.0, 1.0, 0.0}
	my_camera.top = Vector{0.0, 0.0, 1.0}
	//End testing sprites

	physics_objects := new(list.List)
	physics_objects.PushBack(my_camera)

	dt, err := time.ParseDuration("16.67ms")
	if err != nil {
		die(err)
	}

	fmt.Println("Timestep: ", dt)
	var t float64 = 0
	running := true
	for running {
		t0 := time.Now()
		update(physics_objects, float64(dt.Nanoseconds())/1e6)
		t += 0.025
		new_guy.X = math.Cos(t)
		new_guy.Y = math.Sin(t) + 3.0

		//my_camera.point_at(Vector{new_guy.X, new_guy.Y, new_guy.Z})

		general_render(renderees, my_camera)
		glfw.SwapBuffers()

		dt = time.Since(t0)
		rate := time.Microsecond * 16670
		if dt < rate {
			time.Sleep((rate - dt))
			//fmt.Println("This frame took", float64(dt.Nanoseconds())/1000000, "ms to render.")
		}
		running = glfw.Key(glfw.KeyEsc) == 0 && glfw.WindowParam(glfw.Opened) != 0
	}
}

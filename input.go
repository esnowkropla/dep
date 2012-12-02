package main

import (
	"github.com/jteeuwen/glfw"
	"fmt"
)

func print_top() {
	fmt.Printf("Top: %0.3f %0.3f %0.3f\n", my_camera.top.x, my_camera.top.y, my_camera.top.z)
	fmt.Printf("Camera: %0.3f %0.3f %0.3f\n", my_camera.top.x, my_camera.top.y, my_camera.top.z)
}

func up_arrow (press int) {
	fmt.Println("Key Up", press)
	fmt.Println(my_camera.pos)
	if press == 1 {
		my_camera.vel = my_camera.vel.add(vector{0, 0, 1})
		fmt.Println(my_camera.vel)
	} else {
		my_camera.vel = my_camera.vel.add(vector{0, 0, -1})
		fmt.Println(my_camera.vel)
	}
}

func down_arrow (press int) {
	fmt.Println("Key Down", press)
	if press == 1 {
		my_camera.vel = my_camera.vel.add(vector{0, 0, -1})
	} else {
		my_camera.vel = my_camera.vel.add(vector{0, 0, 1})
	}
}

func left_arrow (press int) {
	fmt.Println("Key Left", press)
	if press == 1 {
		my_camera.vel = my_camera.vel.add(vector{-1, 0, 0})
	} else {
		my_camera.vel = my_camera.vel.add(vector{1, 0, 0})
	}
}

func right_arrow (press int) {
	fmt.Println("Key Right", press)
	if press == 1 {
		my_camera.vel = my_camera.vel.add(vector{1, 0, 0})
	} else {
		my_camera.vel = my_camera.vel.add(vector{-1, 0, 0})
	}
}

func key_callback(char, press int) {
	switch char {
		case glfw.KeyUp:
			up_arrow(press)
		case glfw.KeyDown:
			down_arrow(press)
		case glfw.KeyLeft:
			left_arrow(press)
		case glfw.KeyRight:
			right_arrow(press)
		default:
			fmt.Println("None of the above", char)
	}
}

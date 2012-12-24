package main

import (
	"fmt"
	"github.com/jteeuwen/glfw"
	. "vector"
)

func print_top() {
	fmt.Printf("Top: %0.3f %0.3f %0.3f\n", my_camera.top.X, my_camera.top.Y, my_camera.top.Z)
	fmt.Printf("Camera: %0.3f %0.3f %0.3f\n", my_camera.top.X, my_camera.top.Y, my_camera.top.Z)
}

func up_arrow(press int) {
	fmt.Println("Key Up", press)
	fmt.Println(my_camera.pos)
	if press == 1 {
		my_camera.input = my_camera.input.Add(Vector{0, 0, 1})
		fmt.Println(my_camera.input)
	} else {
		my_camera.input = my_camera.input.Add(Vector{0, 0, -1})
		fmt.Println(my_camera.input)
	}
}

func down_arrow(press int) {
	fmt.Println("Key Down", press)
	if press == 1 {
		my_camera.input = my_camera.input.Add(Vector{0, 0, -1})
	} else {
		my_camera.input = my_camera.input.Add(Vector{0, 0, 1})
	}
}

func left_arrow(press int) {
	fmt.Println("Key Left", press)
	if press == 1 {
		my_camera.input = my_camera.input.Add(Vector{-1, 0, 0})
	} else {
		my_camera.input = my_camera.input.Add(Vector{1, 0, 0})
	}
}

func right_arrow(press int) {
	fmt.Println("Key Right", press)
	if press == 1 {
		my_camera.input = my_camera.input.Add(Vector{1, 0, 0})
	} else {
		my_camera.input = my_camera.input.Add(Vector{-1, 0, 0})
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

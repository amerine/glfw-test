package main

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}
	defer glfw.Terminate()

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 2)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	window, err := glfw.CreateWindow(400, 400, "glfw testing", nil, nil)
	if err != nil {
		log.Fatalln("failed to create glfw window:", err)
	}

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

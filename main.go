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

	window, err := glfw.CreateWindow(400, 400, "glfw testing", nil, nil)
	if err != nil {
		log.Fatalln("failed to create glfw window:", err)
	}

	window.SetSizeLimits(400, 400, 400, 400)
	window.MakeContextCurrent()

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

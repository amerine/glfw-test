package main

import (
	"fmt"
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.2/glfw"
)

const (
	windowHeight = 400
	windowWidth  = 400
)

func init() {
	runtime.LockOSThread()
}

func main() {
	window := initGLFWWindow()
	defer glfw.Terminate()

	if err := gl.Init(); err != nil {
		log.Fatalln("failed to initialize OpenGL:", err)
	}

	openGLVersion := gl.GoStr(gl.GetString(gl.VERSION))
	fmt.Println("OpenGL Version:", openGLVersion)

	for !window.ShouldClose() {
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

func initGLFWWindow() *glfw.Window {
	err := glfw.Init()
	if err != nil {
		log.Fatalln("failed to initialize glfw:", err)
	}

	glfw.WindowHint(glfw.Resizable, glfw.False)
	glfw.WindowHint(glfw.ContextVersionMajor, 4)
	glfw.WindowHint(glfw.ContextVersionMinor, 1)
	glfw.WindowHint(glfw.OpenGLProfile, glfw.OpenGLCoreProfile)
	glfw.WindowHint(glfw.OpenGLForwardCompatible, glfw.True)

	window, err := glfw.CreateWindow(windowWidth, windowHeight, "glfw testing", nil, nil)
	if err != nil {
		log.Fatalln("failed to create glfw window:", err)
	}

	window.MakeContextCurrent()

	return window
}

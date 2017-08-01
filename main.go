package main

import (
	"runtime"

	"github.com/go-gl/glfw/v3.2/glfw"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	err := glfw.Init()
	if err != nil (
		log.Fatalln("failed to initialize glfw:", err)
	)
}

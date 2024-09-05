package emulator

import (
	"log"
	"runtime"

	"github.com/go-gl/glfw/v3.3/glfw"
)

const (
	width  = 64
	height = 32
	scale  = 8
	title  = "CHIP-8"
)

func init() {
	runtime.LockOSThread()
}

func Run() {
	err := glfw.Init()
	if err != nil {
		log.Fatalln(err)
	}
	defer glfw.Terminate()

	window, err := glfw.CreateWindow(width*scale, height*scale, title, nil, nil)
	if err != nil {
		log.Fatalln(err)
	}

	window.MakeContextCurrent()

	for !window.ShouldClose() {
		// Do OpenGL stuff.
		window.SwapBuffers()
		glfw.PollEvents()
	}
}

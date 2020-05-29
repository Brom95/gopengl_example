package main

import (
	"fmt"

	// "../gopengl"
	"github.com/Brom95/gopengl"
)

func main() {
	gopengl.GlfvInit()
	gopengl.GlfvWindowHint(gopengl.GLFW_SAMPLES, 4)
	gopengl.GlfvWindowHint(gopengl.GLFW_CONTEXT_VERSION_MAJOR, 3)
	gopengl.GlfvWindowHint(gopengl.GLFW_CONTEXT_VERSION_MINOR, 3)
	gopengl.GlfvWindowHint(gopengl.GLFW_OPENGL_FORWARD_COMPAT, gopengl.GL_TRUE)           // Для того, что бы сделать MacOS счастливой; может не понадобиться
	gopengl.GlfvWindowHint(gopengl.GLFW_OPENGL_PROFILE, gopengl.GLFW_OPENGL_CORE_PROFILE) // Мы не хотим старый OpenGL
	window := gopengl.Window
	window, err := gopengl.GlfvCreateWindow(1024, 768, "Tutorial 01", (*gopengl.GLFWmonitor)(gopengl.NULL), (*gopengl.GLFWwindow)(gopengl.NULL))
	if err != nil {
		fmt.Println("Failed to open GLFW window. If you have an Intel GPU, they are not 3.3 compatible. Try the 2.1 version of the tutorials.")
		gopengl.GlfvTerminate()
	}
	gopengl.GlfvMakeContextCurrent(window) // Инициализируем GLEW
	gopengl.GlewInit()
	gopengl.GlfvSetInputMode(window, gopengl.GLFW_STICKY_KEYS, gopengl.GL_TRUE)
	bufferFree := func() {
		gopengl.GlfwSwapBuffers(window)
		gopengl.GlfwPollEvents()
	}
	bufferFree()
	for gopengl.GlfwGetKey(window, gopengl.GLFW_KEY_ESCAPE) != gopengl.GLFW_PRESS &&
		gopengl.GlfwWindowShouldClose(window) == 0 {
		bufferFree()
	}
}

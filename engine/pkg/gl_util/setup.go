package gl_util

import (
	"log"
	"runtime"

	"github.com/go-gl/gl/v4.1-core/gl"
	"github.com/go-gl/glfw/v3.3/glfw"
)

func init() {
	log.Println("locking glfw goroutine to the main thread")
	// GLFW event handling must run on the main OS thread
	runtime.LockOSThread()
}

func UnInitialize() {
	log.Println("terminating glfw context")
	glfw.Terminate()
	log.Println("unlocking os thread")
	runtime.UnlockOSThread()
}

type GlfwConfig struct {
	windowHints               map[glfw.Hint]int
	fullscreen                bool
	windowWidth, windowHeight uint
	windowTitle               string
}

func (ic *GlfwConfig) Hint(hint glfw.Hint, value int) *GlfwConfig {
	ic.windowHints[hint] = value
	return ic
}

func (ic *GlfwConfig) WindowSize(width, height uint) *GlfwConfig {
	ic.windowHeight = height
	ic.windowWidth = width
	return ic
}

func (ic *GlfwConfig) Fullscreen(fullscreen bool) *GlfwConfig {
	ic.fullscreen = fullscreen
	return ic
}

func (ic *GlfwConfig) Title(title string) *GlfwConfig {
	ic.windowTitle = title
	return ic
}

func DefaultGlfwConfig() *GlfwConfig {
	return &GlfwConfig{
		windowHints: map[glfw.Hint]int{
			glfw.ContextVersionMajor:     4,
			glfw.ContextVersionMinor:     1,
			glfw.OpenGLProfile:           glfw.OpenGLCoreProfile,
			glfw.OpenGLForwardCompatible: glfw.True,
			glfw.Resizable:               glfw.False,
		},
		fullscreen:   false,
		windowWidth:  800,
		windowHeight: 600,
		windowTitle:  "Open GL!",
	}
}

func InitializeGlfwWindow(config *GlfwConfig) (*glfw.Window, error) {
	// Initialize the GLFW context
	log.Println("initializing GLFW")
	if glfwErr := glfw.Init(); glfwErr != nil {
		log.Println("failed to initialize glfw:", glfwErr)
		return nil, glfwErr
	}

	configureWindowHints(config)

	window, windowErr := createWindow(config)
	if windowErr != nil {
		return nil, windowErr
	}
	if glowErr := RebindGlow(); glowErr != nil {
		return nil, glowErr
	}

	return window, nil
}

func RebindGlow() error {
	// Initialize Glow
	// This will load all the function pointers needed to be graphics context agnostic. Or close to it. This must be
	// called after every context switch. We will only have 1 context, so it should be ok though for now.
	log.Println("initializing GLOW")
	if glowErr := gl.Init(); glowErr != nil {
		log.Println("failed to initialize glow bindings:", glowErr)
		return glowErr
	}
	return nil
}

func createWindow(config *GlfwConfig) (*glfw.Window, error) {
	// create the window
	log.Println("initializing window")
	var window *glfw.Window
	var windowErr error
	if config.fullscreen {
		window, windowErr = glfw.CreateWindow(int(config.windowWidth), int(config.windowHeight), config.windowTitle, glfw.GetPrimaryMonitor(), nil) // full screen
	} else {
		window, windowErr = glfw.CreateWindow(int(config.windowWidth), int(config.windowHeight), config.windowTitle, nil, nil) // windowed
	}

	if windowErr != nil {
		log.Println("failed to initialize window:", windowErr)
		return nil, windowErr
	} else {
		window.MakeContextCurrent()
		return window, nil
	}
}

func configureWindowHints(config *GlfwConfig) {
	log.Println("configuring GLFW hints")
	for hint, value := range config.windowHints {
		glfw.WindowHint(hint, value)
	}
}

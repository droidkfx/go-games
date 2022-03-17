package renderer

import (
	"errors"
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/go-gl/gl/v4.1-core/gl"
)

const (
	shaderDir = "./assets/shader/"
)

type Shader struct {
	shaderSrc map[uint32]uint32
	program   uint32
	inuse     bool
}

func NewShader(name string) (Shader, error) {
	shaderSrcs, srcErr := loadShaderFiles(name)
	if srcErr != nil {
		return Shader{}, srcErr
	}

	shaderProgram := gl.CreateProgram()
	shaderSrc := make(map[uint32]uint32, 2)

	log.Println("compiling cloaded shaders")
	for shaderType, src := range shaderSrcs {
		shaderPart, compileErr := compileShaderFile(shaderType, src)
		if compileErr != nil {
			return Shader{}, compileErr
		}
		gl.AttachShader(shaderProgram, shaderPart)
		shaderSrc[shaderType] = shaderPart
		gl.DeleteShader(shaderPart)
	}

	log.Println("linking shader shaders")
	if linkErr := linkShader(shaderProgram); linkErr != nil {
		return Shader{}, linkErr
	}

	shader := Shader{program: shaderProgram, shaderSrc: shaderSrc}

	return shader, nil
}

func (s *Shader) Use() {
	if !s.inuse {
		gl.UseProgram(s.program)
		s.inuse = true
	}
}

func (s *Shader) Detach() {
	gl.UseProgram(0)
	s.inuse = false
}

func linkShader(shaderProgram uint32) error {
	gl.LinkProgram(shaderProgram)

	var status int32
	gl.GetProgramiv(shaderProgram, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetProgramiv(shaderProgram, gl.INFO_LOG_LENGTH, &logLength)

		linkLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(shaderProgram, logLength, nil, gl.Str(linkLog))
		return errors.New(fmt.Sprintf("failed to link program\n%s\n", linkLog))
	}
	return nil
}

func compileShaderFile(shaderType uint32, dat []byte) (uint32, error) {
	shader := gl.CreateShader(shaderType)
	cSource, free := gl.Strs(string(dat))
	defer free()
	gl.ShaderSource(shader, 1, cSource, nil)
	gl.CompileShader(shader)

	var status int32
	gl.GetShaderiv(shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		var logLength int32
		gl.GetShaderiv(shader, gl.INFO_LOG_LENGTH, &logLength)

		shaderLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(shader, logLength, nil, gl.Str(shaderLog))
		return 0, errors.New(fmt.Sprintf("failed to compile glsl shader\n%s\n", shaderLog))
	}

	return shader, nil
}

func loadShaderFiles(prefix string) (map[uint32][]byte, error) {
	log.Printf("loading shader files named '%s'\n", prefix)
	files, listErr := ioutil.ReadDir(shaderDir)
	if listErr != nil {
		return nil, listErr
	}

	mappedFiles := make(map[uint32][]byte, 2)
	for _, v := range files {
		if !v.IsDir() {
			parts := strings.Split(v.Name(), ".")
			if len(parts) == 3 && parts[0] == prefix && parts[2] == "glsl" {
				switch parts[1] {
				case "frag":
					dat, loadErr := loadFile(v)
					if loadErr != nil {
						return nil, loadErr
					}
					mappedFiles[gl.FRAGMENT_SHADER] = dat
				case "vert":
					dat, loadErr := loadFile(v)
					if loadErr != nil {
						return nil, loadErr
					}
					mappedFiles[gl.VERTEX_SHADER] = dat
				default:
					return nil, errors.New(fmt.Sprintf("unexpected shader type %s", parts[1]))
				}

			}
		}
	}

	return mappedFiles, nil
}

func loadFile(v fs.FileInfo) ([]byte, error) {
	dat, readErr := os.ReadFile(shaderDir + v.Name())
	if readErr != nil {
		log.Printf("failed to load shader \n\t%e\n", readErr)
		return nil, readErr
	}
	// c strings end in 0x00
	dat = append(dat, 0)
	return dat, nil
}

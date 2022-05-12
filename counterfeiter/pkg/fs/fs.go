package fs

import "io/ioutil"

var Dependency dependency

type dependency struct {
	ReadFile func(filename string) ([]byte, error)
}

func init() {
	Default()
}

func Default() {
	Dependency.ReadFile = ioutil.ReadFile
}

//go:generate go run github.com/maxbrunsfeld/counterfeiter/v6 -generate

//counterfeiter:generate . FS
type FS interface {
	ReadFile(filename string) ([]byte, error)
}

type filesystem struct{}

func New() FS {
	return &filesystem{}
}

func (fs filesystem) ReadFile(filename string) ([]byte, error) {
	return Dependency.ReadFile(filename)
}

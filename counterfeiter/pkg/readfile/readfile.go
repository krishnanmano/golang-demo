package readfile

import (
	"example.local/pkg/fs"
	"example.local/pkg/logger"
)

var (
	Dependency dependency
)

type dependency struct {
	FS fs.FS
}

func init() {
	Default()
}

func Default() {
	Dependency.FS = fs.New()
}

func Run() error {
	contents, err := Dependency.FS.ReadFile("demo.txt")
	if err != nil {
		return err
	}

	logger.Info(string(contents))
	return nil
}

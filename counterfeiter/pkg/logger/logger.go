package logger

import (
	"io"
	"log"
	"os"
)

func init() {
	Default()
}

func Default() {
	SetOutput(os.Stdout)
}

func Info(txt string) {
	log.Println(txt)
}

func Error(txt string) {
	log.Println(txt)
}

func SetOutput(writer io.Writer) {
	log.SetOutput(writer)
}

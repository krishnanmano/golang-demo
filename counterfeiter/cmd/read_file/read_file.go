package main

import (
	"example.local/pkg/logger"
	"example.local/pkg/readfile"
	"os"
)

func main() {
	err := readfile.Run()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	logger.Info("Read file Complete!")
	os.Exit(0)
}

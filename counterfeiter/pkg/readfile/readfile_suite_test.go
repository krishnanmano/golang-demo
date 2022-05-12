package readfile_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReadfile(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Readfile Suite")
}

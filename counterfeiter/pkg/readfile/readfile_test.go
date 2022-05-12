package readfile_test

import (
	"errors"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"

	"example.local/pkg/fs/fsfakes"
	"example.local/pkg/logger"
	"example.local/pkg/readfile"
)

var _ = Describe("Read File tests", func() {
	var fakeFS *fsfakes.FakeFS
	var logbuf *Buffer

	BeforeEach(func() {
		fakeFS = &fsfakes.FakeFS{}
		readfile.Dependency.FS = fakeFS

		logbuf = NewBuffer()
		logger.SetOutput(logbuf)
	})

	When("Read File successful", func() {
		It("Should not return an error", func() {
			fakeFS.ReadFileReturns([]byte("counterfeiter"), nil)

			err := readfile.Run()

			Expect(err).To(BeNil())
			Expect(logbuf).To(Say("counterfeiter"))
			Expect(fakeFS.ReadFileArgsForCall(0)).To(Equal("demo.txt"))
		})
	})

	When("Read File Failed to read demo.txt", func() {
		It("Should return an error", func() {
			fakeFS.ReadFileReturns(nil, errors.New("failed to read demo.txt"))

			err := readfile.Run()

			Expect(err).To(MatchError("failed to read demo.txt"))
		})
	})
})

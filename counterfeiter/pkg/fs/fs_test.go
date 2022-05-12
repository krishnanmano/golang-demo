package fs_test

import (
	"errors"
	"example.local/pkg/fs"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("FS tests", func() {
	var readFileParams []string
	var readFileCallCount int

	BeforeEach(func() {
		readFileParams = make([]string, 0, 0)
		readFileCallCount = 0
		fs.Dependency.ReadFile = func(filename string) ([]byte, error) {
			readFileCallCount++
			readFileParams = append(readFileParams, filename)
			return []byte("counterfeiter"), nil
		}
	})

	Context("ReadFile functionality tests", func() {
		When("ReadFile() successful in read the given file", func() {
			It("should not report an error", func() {
				f := fs.New()
				contents, err := f.ReadFile("demo.txt")

				Expect(err).To(BeNil())
				Expect(string(contents)).To(Equal("counterfeiter"))
				Expect(readFileParams[0]).To(Equal("demo.txt"))
				Expect(readFileCallCount).To(Equal(1))
			})
		})

		When("ReadFile() fails to read file", func() {
			It("should report an error", func() {
				fs.Dependency.ReadFile = func(filename string) ([]byte, error) {
					readFileCallCount++
					readFileParams = append(readFileParams, filename)
					return nil, errors.New("failed to read demo.txt")
				}

				f := fs.New()
				contents, err := f.ReadFile("demo.txt")

				Expect(err).To(MatchError("failed to read demo.txt"))
				Expect(contents).To(BeNil())
				Expect(readFileParams[0]).To(Equal("demo.txt"))
				Expect(readFileCallCount).To(Equal(1))
			})
		})
	})
})

package main

type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}

type readWriter struct{}

func (rw readWriter) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (rw readWriter) Write(p []byte) (n int, err error) {
	return 0, nil
}

func main() {
	var rw readWriter = readWriter{}
	var r []byte
	rw.Read(r)

	w := []byte("Jerry")
	rw.Write(w)
}

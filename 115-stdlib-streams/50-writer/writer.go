// Writing files in Go follows similar patterns to the
// ones we saw earlier for reading.

package main

import (
	"os"
)

type MyWriter struct{}

func (writer *MyWriter) Write(p []byte) (n int, err error) {
	prefix := []byte("[BOO] ")
	data := append(prefix, p...)
	os.Stdout.Write(data)
	return len(p) + len(prefix), nil
}

func main() {
	str := []byte("Hoł hoł hoł\n")

	writer := os.Stdout
	writer.Write(str)

	superWriter := MyWriter{}
	superWriter.Write(str)
}

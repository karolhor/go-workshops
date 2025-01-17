package main

import (
	"io"
	"os"
	"os/exec"
)

func main() {

	c1 := exec.Command("ls", "/etc")
	c2 := exec.Command("wc", "-l")

	r, w := io.Pipe()

	c1.Stdout = w
	c2.Stdin = r

	c2.Stdout = os.Stdout

	c1.Start()
	c2.Start()
	c1.Wait()
	w.Close()
	c2.Wait()
}

package main

import "syscall"
import "os"
import "os/exec"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	binary, err := exec.LookPath("ls")
	check(err)

	args := []string{"ls", "-a", "-l", "-h"}
	env := os.Environ()

	execErr := syscall.Exec(binary, args, env)
	check(execErr)
}

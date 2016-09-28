package main

import "fmt"
import "io/ioutil"
import "os/exec"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	dateCmd := exec.Command("date")

	dateOut, err := dateCmd.Output()
	check(err)

	fmt.Println("> date")
	fmt.Println(string(dateOut))

	grepCmd := exec.Command("grep", "hello")

	grepIn, err := grepCmd.StdinPipe()
	check(err)

	grepOut, err := grepCmd.StdoutPipe()
	check(err)

	grepCmd.Start()

	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()

	grepBytes, err := ioutil.ReadAll(grepOut)
	check(err)

	grepCmd.Wait()

	fmt.Println("> grep hello")
	fmt.Println(string(grepBytes))

	lsCmd := exec.Command("bash", "-c", "ls -a -l -h")
	lsOut, err := lsCmd.Output()
	check(err)

	fmt.Println("> ls -a -l -h")
	fmt.Println(string(lsOut))
}

package main

import (
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	cmd := exec.Command("go", "-c", "echo stdout; echo 1>&2 stderr")
	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Fatal(err)
	}

	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}

	slurp, _ := ioutil.ReadAll(stderr)
	log.Printf("%s\n", slurp)

	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
}

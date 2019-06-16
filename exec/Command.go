package main

import (
	"bytes"
	"log"
	"os/exec"
	"strings"
)

func main() {
	cmd := exec.Command("npm", "a-z", "A-Z")
	cmd.Stdin = strings.NewReader("some input")

	var out bytes.Buffer

	cmd.Stdout = &out
	err := cmd.Run()

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("in all caps: %q\n", out.String())
}

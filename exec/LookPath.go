package main

import (
	"log"
	"os/exec"
)

func main() {
	path, err := exec.LookPath("npm")

	if err != nil {
		log.Fatal("installing fortune is in your future")
	}

	log.Printf("fortune is available at %s\n", path)
}

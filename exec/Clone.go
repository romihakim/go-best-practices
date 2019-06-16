package main

import (
	"bytes"
	// "encoding/json"
	// "os"
	"os/exec"

	tk "github.com/eaciit/toolkit"
)

func main() {
	cmd := exec.Command("git", "clone", "https://github.com/kodesain/dashgrid.js", "backup")
	var out bytes.Buffer
	cmd.Stdout = &out
	tk.Println(out)
	err := cmd.Run()
	if err != nil {
		tk.Println(err.Error())
	} else {
		tk.Println(out.String())
	}
}

// package main

// import (
// 	"log"
// 	"os"
// 	"os/exec"
// )

// func main() {
// 	cmd := exec.Command("git", "clone", "https://github.com/romiamirul/golang.git", "repo")
// 	err := cmd.Run()

// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }

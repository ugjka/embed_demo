package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

const template = `package main

var image = []byte("%s")
`

const filename = "data.go"

func main() {
	log.Println("generating " + filename)
	// slurp the file
	data, err := ioutil.ReadFile("kitty.png")
	if err != nil {
		log.Fatal(err)
	}
	// bytes.Buffer is faster than string operations
	buf := bytes.NewBuffer(nil)
	for _, v := range data {
		// write escaped hex value of each byte into the buffer
		fmt.Fprintf(buf, `\x%02X`, v)
	}
	// fill the template
	out := fmt.Sprintf(template, buf)
	// remove existing file if it exists
	os.Remove(filename)
	err = ioutil.WriteFile(filename, []byte(out), 0644)
	if err != nil {
		log.Fatal(err)
	}
	// format our generated go file for a good measure
	exec.Command("gofmt", "-w", filename).Run()
}

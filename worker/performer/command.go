package performer

import (
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

type Input struct {
	App  string
	Args []string
	Env  []string
}

func Standard(input Input) error {
	var out string

	cmd := exec.Command(input.App, input.Args...)
	cmd.Env = append(os.Environ(), input.Env...)

	stdout, errSo := cmd.StdoutPipe()
	if errSo != nil {
		log.Fatal(errSo)
	}

	stderr, errSe := cmd.StderrPipe()
	if errSe != nil {
		log.Fatal(errSe)
	}

	errStt := cmd.Start()
	if errStt != nil {
		b, _ := ioutil.ReadAll(stderr)
		out = string(b)
		// stdoutReader.ReadRune()
	} else {
		b, _ := ioutil.ReadAll(stdout)
		out = string(b)
	}

	if errWt := cmd.Wait(); errWt != nil {
		log.Fatal(errWt)
	}

	log.Printf("command out: %v", out)

	return errStt
}

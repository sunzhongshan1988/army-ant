package performer

import (
	"io"
	"log"
	"os"
	"os/exec"
)

type Input struct {
	App  string
	Args []string
	Env  []string
}

type Output struct {
	StdoutPipeOut io.ReadCloser
	StdoutPipeErr error
}

func Standard(input Input) Output {
	cmd := exec.Command(input.App, input.Args...)
	cmd.Env = append(os.Environ(), input.Env...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	return Output{
		StdoutPipeOut: stdout,
		StdoutPipeErr: err,
	}
}

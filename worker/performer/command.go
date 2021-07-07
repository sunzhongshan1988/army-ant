package performer

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/sunzhongshan1988/army-ant/worker/grpc"
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

func Standard(input Input) {
	var out string
	var status int32
	startAt := ptypes.TimestampNow()

	cmd := exec.Command(input.App, input.Args...)
	cmd.Env = append(os.Environ(), input.Env...)

	stdout, errSo := cmd.StdoutPipe()
	if errSo != nil {
		log.Printf("[command,stdoutpipe] error: %s", errSo)
	}

	stderr, errSe := cmd.StderrPipe()
	if errSe != nil {
		log.Printf("[command,stderrpipe] error: %s", errSe)
	}

	errStt := cmd.Start()
	if errStt != nil {
		b, _ := ioutil.ReadAll(stderr)
		status = 1
		out = string(b)
		// stdoutReader.ReadRune()
	} else {
		b, _ := ioutil.ReadAll(stdout)
		status = 0
		out = string(b)
	}

	if errWt := cmd.Wait(); errWt != nil {
		log.Printf("[command,wait] error: %s", errWt)
	}

	log.Printf("[command,wait] info: %v", out)

	grpc.TaskResult(out, status, startAt, ptypes.TimestampNow())
}

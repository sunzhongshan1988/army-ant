package performer

import (
	"github.com/golang/protobuf/ptypes"
	"github.com/sunzhongshan1988/army-ant/worker/grpc"
	"github.com/sunzhongshan1988/army-ant/worker/model"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

func Standard(input model.Input) {
	commandResult := &model.CommandResult{
		TaskID:     input.TaskID,
		TaskName:   input.TaskName,
		TaskRemark: input.TaskRemark,
		InstanceID: input.InstanceID,
		Out:        "",
		Type:       input.Type,
		Status:     0,
		StartAt:    ptypes.TimestampNow(),
		EndAt:      nil,
	}

	cmd := exec.Command(input.App, input.Args...)
	cmd.Env = append(os.Environ(), input.Env...)
	cmd.Dir = input.Dir

	stdout, errSo := cmd.StdoutPipe()
	if errSo != nil {
		commandResult.Status = 1
		log.Printf("[command,stdoutpipe] error: %s", errSo)
	}

	stderr, errSe := cmd.StderrPipe()
	if errSe != nil {
		commandResult.Status = 1
		log.Printf("[command,stderrpipe] error: %s", errSe)
	}

	errStt := cmd.Start()
	if errStt != nil {
		b, _ := ioutil.ReadAll(stderr)
		commandResult.Status = 1
		commandResult.Out = string(b)
		// stdoutReader.ReadRune()
	} else {
		b, _ := ioutil.ReadAll(stdout)
		commandResult.Status = 0
		commandResult.Out = string(b)
	}

	if errWt := cmd.Wait(); errWt != nil {
		commandResult.Status = 1
		log.Printf("[command,wait] error: %s", errWt)
	}

	log.Printf("[command,out] info: %v", commandResult.Out)

	commandResult.EndAt = ptypes.TimestampNow()
	grpc.TaskResult(commandResult)
}

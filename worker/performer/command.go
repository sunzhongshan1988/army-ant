package performer

import (
	"github.com/sunzhongshan1988/army-ant/worker/grpc"
	"github.com/sunzhongshan1988/army-ant/worker/model"
	"google.golang.org/protobuf/types/known/timestamppb"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var execInstance = make(map[string]*exec.Cmd)

func Standard(input model.Input) {
	commandResult := &model.CommandResult{
		TaskID:     input.TaskID,
		TaskName:   input.TaskName,
		TaskRemark: input.TaskRemark,
		InstanceID: input.InstanceID,
		Error:      "",
		Output:     "",
		Type:       input.Type,
		Status:     0,
		StartAt:    timestamppb.Now(),
		EndAt:      nil,
	}

	execInstance[input.TaskID] = exec.Command(input.App, input.Args...)
	execInstance[input.TaskID].Env = append(os.Environ(), input.Env...)
	execInstance[input.TaskID].Dir = input.Dir

	stdout, errSo := execInstance[input.TaskID].StdoutPipe()
	if errSo != nil {
		commandResult.Status = 1
		commandResult.Error = errSo.Error()
		log.Printf("[command,stdoutpipe] error: %s", errSo)
	}

	stderr, errSe := execInstance[input.TaskID].StderrPipe()
	if errSe != nil {
		commandResult.Status = 1
		commandResult.Error = errSe.Error()
		log.Printf("[command,stderrpipe] error: %s", errSe)
	}

	errStt := execInstance[input.TaskID].Start()
	if errStt != nil {
		b, _ := ioutil.ReadAll(stderr)
		commandResult.Status = 1
		commandResult.Error = string(b)
		// stdoutReader.ReadRune()
	} else {
		b, _ := ioutil.ReadAll(stdout)
		commandResult.Status = 0
		commandResult.Output = string(b)
	}

	if errWt := execInstance[input.TaskID].Wait(); errWt != nil {
		commandResult.Status = 1
		commandResult.Error = errWt.Error()
		log.Printf("[command,wait] error: %s", errWt)
	}

	delete(execInstance, input.TaskID)

	log.Printf("[command,out] info: %v", commandResult.Output)

	commandResult.EndAt = timestamppb.Now()
	grpc.TaskResult(commandResult)
}

func Kill(taskID string) bool {
	if cmd, ok := execInstance[taskID]; ok {
		err := cmd.Process.Kill()
		if err != nil {
			log.Printf("[command,kill] info: %v", err)
			return false
		}
		log.Printf("[command,kill] info: %v", "Successfully killed task")
		return true
	} else {
		log.Printf("[command,kill] info: %v", "taskID not found")
		return false
	}
}

package model

type Command struct {
	App  string   `json:"app"`
	Args []string `json:"args"`
	Env  []string `json:"env"`
}
type DNA struct {
	Cmd     *Command `json:"cmd"`
	Version string   `json:"version"`
}

type CommandMutation struct {
	Args []string `json:"args"`
	Env  []string `json:"env"`
}
type Mutation struct {
	Cmd     *CommandMutation `json:"cmd"`
	Version string           `json:"version"`
}

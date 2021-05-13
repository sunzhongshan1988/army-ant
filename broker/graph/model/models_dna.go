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

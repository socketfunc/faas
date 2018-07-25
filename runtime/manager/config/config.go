package config

type Runtime struct {
	Port         int    `yaml:"port"`
	InternalPort int    `yaml:"internalPort"`
	ExecTimeout  int    `yaml:"execTimeout"`
	CodePath     string `yaml:"codePath"`
}

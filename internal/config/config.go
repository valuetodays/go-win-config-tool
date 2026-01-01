package config

type EnvConfig struct {
	Name   string   `yaml:"name"`
	Value  string   `yaml:"value,omitempty"`
	Append []string `yaml:"append,omitempty"`
	Scope  string   `yaml:"scope"` // user | system
}
type RootConfig struct {
	Root struct {
		Paths []string    `yaml:"paths"`
		Envs  []EnvConfig `yaml:"envs"`
	} `yaml:"root"`
}

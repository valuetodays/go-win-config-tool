package config

type EnvConfig struct {
	Name   string   `yaml:"name"`
	Value  string   `yaml:"value,omitempty"`
	Append []string `yaml:"append,omitempty"`
	Scope  string   `yaml:"scope"` // user | system
}

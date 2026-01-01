package config

type EnvConfig struct {
	Name   string   `yaml:"name"`
	Value  string   `yaml:"value,omitempty"`
	Append []string `yaml:"append,omitempty"`
	Scope  string   `yaml:"scope"` // user | system
}
type SoftwareConfig struct {
	Name    string `yaml:"name"`
	RootDir string `yaml:"root_dir"`
}
type RootConfig struct {
	Root struct {
		Paths     []string         `yaml:"paths"`
		Envs      []EnvConfig      `yaml:"envs"`
		Softwares []SoftwareConfig `yaml:"softwares"`
	} `yaml:"root"`
}

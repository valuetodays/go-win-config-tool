package config

type RootConfig struct {
	Root struct {
		Paths []string `yaml:"paths"`
	} `yaml:"root"`
}

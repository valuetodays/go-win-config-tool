package config

type ShortcutConfig struct {
	Name   string `yaml:"name"`
	Source string `yaml:"source"`
	Target string `yaml:"target"`
	GenCmd string `yaml:"gencmd"`
}

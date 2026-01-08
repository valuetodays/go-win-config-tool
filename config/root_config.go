package config

type RootConfig struct {
	Root struct {
		DefDirs   DefDirsConfig    `yaml:"def_dirs"`
		Paths     []string         `yaml:"paths"`
		Envs      []EnvConfig      `yaml:"envs"`
		Softwares []SoftwareConfig `yaml:"softwares"`
		Shortcuts []ShortcutConfig `yaml:"shortcuts"`
	} `yaml:"root"`
}

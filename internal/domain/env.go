package domain

type EnvStatus struct {
	Name    string
	Scope   string
	Exists  bool
	Correct bool
	Missing []string
	Value   string
}

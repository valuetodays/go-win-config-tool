package domain

type EnvStatus struct {
	Name    string   `json:"name"`
	Scope   string   `json:"scope"`
	Exists  bool     `json:"exists"`
	Correct bool     `json:"correct"`
	Missing []string `json:"missing"`
	Value   string   `json:"value"`
}

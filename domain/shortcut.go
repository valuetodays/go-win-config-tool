package domain

type ShortcutStatus struct {
	Name    string `json:"name"`
	Path    string `json:"path"`
	Target  string `json:"target"`
	Exists  bool   `json:"exists"`
	Correct bool   `json:"correct"`
}

package dto

type ShortcutStatus struct {
	Name    string `json:"name"`
	Source  string `json:"source"`
	Target  string `json:"target"`
	Exists  bool   `json:"exists"`
	Correct bool   `json:"correct"`
}

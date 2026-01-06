package domain

type PathStatus struct {
	Path   string `json:"path"`
	Exists bool   `json:"exists"`
}

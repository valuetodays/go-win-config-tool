package dto

type SoftwareStatus struct {
	Name    string `json:"name"`
	RootDir string `json:"rootDir"`
	Exists  bool   `json:"exists"`
}

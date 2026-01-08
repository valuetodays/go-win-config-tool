package service

import (
	"go-win-config-tool/internal/dto"
	"os"
	"path/filepath"
	"strings"
)

type PathService struct {
	paths []string
}

func NewPathService(paths []string) *PathService {
	return &PathService{paths: paths}
}

func expandPath(p string) string {
	if strings.HasPrefix(p, "~") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, p[1:])
	}
	return p
}

func (s *PathService) GetPathsStatus() []dto.PathStatus {
	result := []dto.PathStatus{}

	for _, p := range s.paths {
		realPath := expandPath(p)
		_, err := os.Stat(realPath)

		result = append(result, dto.PathStatus{
			Path:   p,
			Exists: err == nil,
		})
	}
	return result
}

func (s *PathService) CreatePath(path string) error {
	realPath := expandPath(path)
	return os.MkdirAll(realPath, os.ModePerm)
}

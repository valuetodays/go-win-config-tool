package service

import (
	"go-win-config-tool/config"
	"go-win-config-tool/domain"
	"os"
)

type SoftwareService struct {
	softwares []config.SoftwareConfig
}

func NewSoftwareService(sw []config.SoftwareConfig) *SoftwareService {
	return &SoftwareService{softwares: sw}
}

func (s *SoftwareService) GetSoftwareStatus() []domain.SoftwareStatus {
	result := []domain.SoftwareStatus{}

	for _, sw := range s.softwares {
		status := domain.SoftwareStatus{
			Name:    sw.Name,
			RootDir: sw.RootDir,
		}

		if info, err := os.Stat(sw.RootDir); err == nil && info.IsDir() {
			status.Exists = true
		}
		/*
		   for _, item := range sw.Items {
		       full := filepath.Join(sw.RootDir, item.Path)
		       _, err := os.Stat(full)

		       status.Items = append(status.Items, SoftwareItemStatus{
		           Type:   item.Type,
		           Path:   item.Path,
		           Exists: err == nil,
		       })
		   }
		*/
		result = append(result, status)
	}

	return result
}

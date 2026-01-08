package service

import (
	"go-win-config-tool/internal/domain"
	"go-win-config-tool/internal/dto"
	"os"
)

type SoftwareService struct {
	softwares []domain.SoftwareModel
}

func NewSoftwareService(sw []domain.SoftwareModel) *SoftwareService {
	return &SoftwareService{softwares: sw}
}

func (s *SoftwareService) GetSoftwareStatus() []dto.SoftwareStatus {
	result := []dto.SoftwareStatus{}

	for _, sw := range s.softwares {
		status := dto.SoftwareStatus{
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

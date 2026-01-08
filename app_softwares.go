package main

import (
	"go-win-config-tool/internal/dto"
)

func (a *App) GetSoftwareStatus() []dto.SoftwareStatus {
	if a.softwareSvc == nil {
		return nil
	}
	return a.softwareSvc.GetSoftwareStatus()
}

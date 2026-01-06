package main

import (
	"go-win-config-tool/domain"
)

func (a *App) GetSoftwareStatus() []domain.SoftwareStatus {
	if a.softwareSvc == nil {
		return nil
	}
	return a.softwareSvc.GetSoftwareStatus()
}

package main

import (
	"go-win-config-tool/domain"
)

func (a *App) GetPathsStatus() []domain.PathStatus {
	if a.pathSvc == nil {
		return nil
	}
	return a.pathSvc.GetPathsStatus()
}

func (a *App) CreatePath(path string) error {
	if a.pathSvc == nil {
		return nil
	}
	return a.pathSvc.CreatePath(path)
}

package main

import "go-win-config-tool/internal/dto"

func (a *App) GetPathsStatus() []dto.PathStatus {
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

package main

import (
	"go-win-config-tool/domain"
)

func (a *App) GetEnvStatus() ([]domain.EnvStatus, error) {
	if a.envSvc == nil {
		return nil, nil
	}
	return a.envSvc.GetEnvStatus()
}

func (a *App) SetEnv(name, value, scope string) error {
	if a.envSvc == nil {
		return nil
	}
	return a.envSvc.SetEnv(name, value, scope)
}

func (a *App) AppendEnv(name string, values []string, scope string) error {
	if a.envSvc == nil {
		return nil
	}
	return a.envSvc.AppendEnv(name, values, scope)
}

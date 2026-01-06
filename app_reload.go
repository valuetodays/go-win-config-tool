package main

import (
	"go-win-config-tool/config"
	"go-win-config-tool/service"
)

func (a *App) loadConfig(path string) error {
	cfg, err := config.Load(path)
	// cfg, err := config.LoadConfig(path)
	if err != nil {
		return err
	}

	a.configPath = path
	a.cfg = cfg

	a.pathSvc = service.NewPathService(cfg.Root.Paths)
	// a.envSvc = service.NewEnvService(cfg.Root.Envs)
	// a.softwareSvc = service.NewSoftwareService(cfg.Root.Softwares)
	a.shortcutSvc = service.NewShortcutService(cfg.Root.Shortcuts)

	return nil
}

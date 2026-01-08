package main

import (
	"go-win-config-tool/config"
	"go-win-config-tool/internal/domain"
	"go-win-config-tool/internal/util"
	"go-win-config-tool/service"
	"path/filepath"
)

func (a *App) loadConfig(path string) error {
	cfg, err := config.Load(path)
	if err != nil {
		return err
	}

	root := cfg.Root

	// 展开 def_dirs（支持 %USERPROFILE%）
	softwareBase := util.ExpandWinEnv(root.DefDirs.SoftwareDir)
	shortcutBase := util.ExpandWinEnv(root.DefDirs.ShortcutDir)

	softwareModels := []domain.SoftwareModel{}
	// 重新组装 softwares
	for i := range root.Softwares {
		if !filepath.IsAbs(root.Softwares[i].DirName) {
			softwareModel := domain.SoftwareModel{
				Name:    root.Softwares[i].Name,
				RootDir: filepath.Join(softwareBase, root.Softwares[i].DirName),
			}
			softwareModels = append(softwareModels, softwareModel)
		}
	}

	shortcutModels := buildShortcutModels(shortcutBase, root.Shortcuts)

	// 初始化 services（此时配置已经是“最终态”）
	a.pathSvc = service.NewPathService(root.Paths)
	a.envSvc = service.NewEnvService(root.Envs)
	a.softwareSvc = service.NewSoftwareService(softwareModels)
	a.shortcutSvc = service.NewShortcutService(shortcutModels)

	a.cfg = cfg
	a.configPath = path
	return nil
}

func buildShortcutModels(shortcutBase string, shortcutConfigs []config.ShortcutConfig) []domain.ShortcutModel {
	shortcutModels := []domain.ShortcutModel{}
	// 重新组装 shortcuts
	for i := range shortcutConfigs {
		if !filepath.IsAbs(shortcutConfigs[i].Target) {
			shortcutModel := domain.ShortcutModel{
				Name:   shortcutConfigs[i].Name,
				Source: shortcutConfigs[i].Source,
				Target: filepath.Join(shortcutBase, shortcutConfigs[i].Target),
				GenCmd: shortcutConfigs[i].GenCmd,
			}
			shortcutModels = append(shortcutModels, shortcutModel)
		}
	}
	return shortcutModels
}

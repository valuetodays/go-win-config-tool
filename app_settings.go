package main

import (
	"fmt"
	"go-win-config-tool/config"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// func (a *App) LoadConfigFile(path string) error {
// 	return a.loadConfig(path)
// }

func (a *App) SelectAndLoadConfig() (string, error) {
	path, err := runtime.OpenFileDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "选择 config.yml",
		Filters: []runtime.FileFilter{
			{DisplayName: "YAML 文件", Pattern: "*.yml;*.yaml"},
		},
	})
	if err != nil || path == "" {
		return "", err
	}

	fmt.Println("begin. Loading config from:", path)
	err = a.loadConfig(path)
	fmt.Println("end.Loading config from:", path, err)
	// _, err = config.Load(path)
	if err != nil {
		return "", err
	}

	// 🔔 通知前端：配置已变更
	runtime.EventsEmit(a.ctx, "config:changed")

	return path, nil
}

func (a *App) GetCurrentConfigPath() string {
	return config.Path()
}

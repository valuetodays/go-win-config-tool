package main

import (
	"context"
	"fmt"
	"go-win-config-tool/config"
	"go-win-config-tool/domain"
	"go-win-config-tool/service"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx        context.Context
	configPath string
	cfg        *config.RootConfig

	pathSvc *service.PathService
	envSvc  *service.EnvService
	// softwareSvc *service.SoftwareService
	shortcutSvc *service.ShortcutService
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) GetSoftwareStatus() ([]domain.SoftwareStatus, error) {
	result := []domain.SoftwareStatus{}

	for _, sw := range config.Get().Root.Softwares {
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

	return result, nil
}

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

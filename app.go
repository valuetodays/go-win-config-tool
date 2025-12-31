package main

import (
	"context"
	"fmt"
	"go-win-config-tool/internal/config"
	"go-win-config-tool/internal/domain"
	"os"
	"path/filepath"
	"strings"
)

// App struct
type App struct {
	ctx context.Context
	cfg *config.RootConfig
}

// NewApp creates a new App application struct
func NewApp() *App {
	cfg, _ := config.LoadConfig("config.yml")
	return &App{cfg: cfg}
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

func expandPath(p string) string {
	if strings.HasPrefix(p, "~") {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, p[1:])
	}
	return p
}

func (a *App) GetPathsStatus() ([]domain.PathStatus, error) {
	result := []domain.PathStatus{}

	for _, p := range a.cfg.Root.Paths {
		realPath := expandPath(p)
		_, err := os.Stat(realPath)

		result = append(result, domain.PathStatus{
			Path:   p,
			Exists: err == nil,
		})
	}
	return result, nil
}

func (a *App) CreatePath(path string) error {
	realPath := expandPath(path)
	return os.MkdirAll(realPath, os.ModePerm)
}

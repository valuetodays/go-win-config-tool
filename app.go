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

func (a *App) GetEnvStatus() ([]domain.EnvStatus, error) {
	cfg := a.cfg.Root.Envs
	result := []domain.EnvStatus{}

	for _, env := range cfg {
		current := os.Getenv(env.Name)

		status := domain.EnvStatus{
			Name:   env.Name,
			Scope:  env.Scope,
			Exists: current != "",
			Value:  current,
		}

		if env.Value != "" {
			status.Correct = current == env.Value
		}

		if len(env.Append) > 0 {
			for _, p := range env.Append {
				if !strings.Contains(current, p) {
					status.Missing = append(status.Missing, p)
				}
			}
			status.Correct = len(status.Missing) == 0
		}

		result = append(result, status)
	}

	return result, nil
}

func (a *App) SetEnv(name, value, scope string) error {
	return nil
}

func (a *App) AppendEnv(name string, values []string, scope string) error {
	return nil
}

func (a *App) GetSoftwareStatus() ([]domain.SoftwareStatus, error) {
	result := []domain.SoftwareStatus{}

	for _, sw := range a.cfg.Root.Softwares {
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

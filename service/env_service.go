package service

import (
	"go-win-config-tool/config"
	"go-win-config-tool/domain"
	"os"
	"path/filepath"
	"strings"
)

type EnvService struct {
	envs []config.EnvConfig
}

func NewEnvService(envConfig []config.EnvConfig) *EnvService {
	return &EnvService{envs: envConfig}
}

func (a *EnvService) GetEnvStatus() ([]domain.EnvStatus, error) {
	cfg := config.Get().Root.Envs
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
			euqals := strings.EqualFold(
				filepath.Clean(current),
				filepath.Clean(env.Value),
			)
			status.Correct = euqals
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

func (a *EnvService) SetEnv(name, value, scope string) error {
	return nil
}

func (a *EnvService) AppendEnv(name string, values []string, scope string) error {
	return nil
}

package service

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"go-win-config-tool/config"
	"go-win-config-tool/domain"
	"go-win-config-tool/internal/shortcut"
)

type ShortcutService struct {
	shortcuts []config.ShortcutConfig
}

func NewShortcutService(shortcuts []config.ShortcutConfig) *ShortcutService {
	return &ShortcutService{shortcuts: shortcuts}
}

func normalize(p string) string {
	p = filepath.Clean(p)
	return strings.ToLower(p)
}

func (s *ShortcutService) List() ([]domain.ShortcutStatus, error) {
	var result []domain.ShortcutStatus

	// this path should be configured in yml
	desktop := os.Getenv("USERPROFILE") + "\\Desktop\\scs"

	for _, sc := range s.shortcuts {
		path := filepath.Join(desktop, sc.Name)

		status := domain.ShortcutStatus{
			Name:   sc.Name,
			Path:   path,
			Target: sc.Target,
		}

		if _, err := os.Stat(path); err == nil {
			status.Exists = true

			if target, err := shortcut.ReadTarget(path); err == nil {
				status.Correct = normalize(target) == normalize(sc.Target)
			}
		}

		result = append(result, status)
	}

	return result, nil
}

func (s *ShortcutService) Generate(name string) error {
	for _, sc := range s.shortcuts {
		if sc.Name == name {
			cmd := exec.Command("cmd", "/C", sc.GenCmd)
			return cmd.Run()
		}
	}
	return nil
}

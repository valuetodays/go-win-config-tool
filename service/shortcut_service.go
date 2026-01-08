package service

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"go-win-config-tool/internal/domain"
	"go-win-config-tool/internal/dto"
	"go-win-config-tool/internal/shortcut"
	"go-win-config-tool/internal/util"
)

type ShortcutService struct {
	shortcuts []domain.ShortcutModel
}

func NewShortcutService(shortcuts []domain.ShortcutModel) *ShortcutService {
	return &ShortcutService{shortcuts: shortcuts}
}

func normalize(p string) string {
	p = filepath.Clean(p)
	return strings.ToLower(p)
}

func (s *ShortcutService) GetShortcuts() ([]dto.ShortcutStatus, error) {
	var result []dto.ShortcutStatus

	for _, sc := range s.shortcuts {
		path := filepath.Join("", sc.Source)
		// expandedTarget := os.ExpandEnv(sc.Target)
		expandedTarget := util.ExpandWinEnv(sc.Target)
		status := dto.ShortcutStatus{
			Name:   sc.Name,
			Source: path,
			Target: expandedTarget,
		}

		if _, err := os.Stat(path); err == nil {
			status.Exists = true

			if target, err := shortcut.ReadTarget(expandedTarget); err == nil {
				realFileLnkTo := normalize(target)
				sourceFile := normalize(sc.Source)
				status.Correct = (realFileLnkTo == sourceFile)
				status.Source = realFileLnkTo
			}
		}

		result = append(result, status)
	}

	return result, nil
}

func (s *ShortcutService) GenerateShortcut(name string) error {
	for _, sc := range s.shortcuts {
		if sc.Name == name {
			cmd := exec.Command("cmd", "/C", sc.GenCmd)
			return cmd.Run()
		}
	}
	return nil
}

package service

import (
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
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

func expandWinEnv(s string) string {
	re := regexp.MustCompile(`%([^%]+)%`)
	return re.ReplaceAllStringFunc(s, func(match string) string {
		// 去掉 % 符号
		key := match[1 : len(match)-1]
		return os.Getenv(key)
	})
}

func (s *ShortcutService) GetShortcuts() ([]domain.ShortcutStatus, error) {
	var result []domain.ShortcutStatus

	for _, sc := range s.shortcuts {
		path := filepath.Join("", sc.Source)
		// expandedTarget := os.ExpandEnv(sc.Target)
		expandedTarget := expandWinEnv(sc.Target)
		status := domain.ShortcutStatus{
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

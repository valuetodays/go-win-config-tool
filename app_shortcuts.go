package main

import "go-win-config-tool/internal/dto"

func (a *App) GetShortcuts() ([]dto.ShortcutStatus, error) {
	return a.shortcutSvc.GetShortcuts()
}

func (a *App) GenerateShortcut(name string) error {
	return a.shortcutSvc.GenerateShortcut(name)
}

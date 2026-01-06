package main

import "go-win-config-tool/domain"

func (a *App) GetShortcuts() ([]domain.ShortcutStatus, error) {
	return a.shortcutSvc.GetShortcuts()
}

func (a *App) GenerateShortcut(name string) error {
	return a.shortcutSvc.GenerateShortcut(name)
}

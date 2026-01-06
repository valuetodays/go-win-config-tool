package main

func (a *App) LoadConfigFile(path string) error {
	return a.loadConfig(path)
}

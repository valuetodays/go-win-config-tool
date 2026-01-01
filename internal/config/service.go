package config

import (
	"sync"
)

var (
	currentPath string
	currentCfg  *RootConfig
	mu          sync.RWMutex
)

func Load(path string) (*RootConfig, error) {
	cfg, err := LoadConfig(path)
	if err != nil {
		return nil, err
	}

	mu.Lock()
	defer mu.Unlock()
	currentPath = path
	currentCfg = cfg
	return cfg, nil
}

func Get() *RootConfig {
	mu.RLock()
	defer mu.RUnlock()
	return currentCfg
}

func Path() string {
	mu.RLock()
	defer mu.RUnlock()
	return currentPath
}

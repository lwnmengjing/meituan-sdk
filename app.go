package meituan_sdk

import (
	"fmt"
)

// App 应用
type App struct {
	c *Config
}

// NewApp 新建应用
func NewApp(c *Config) (*App, error) {
	if c == nil {
		return nil, fmt.Errorf("app config is nil")
	}
	return &App{c: c}, nil
}

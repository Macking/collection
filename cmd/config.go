package cmd

import (
	"context"
	"github.com/Macking/collection/internal/dbcore"
	"github.com/Macking/collection/internal/miniocore"
	"sync"
)

var globalConfig *Config

type ctxKeyWaitGroup struct{}

type Config struct {
	AppName string
	dbcore.DBConfig
	miniocore.MinioConfig
	Ctx    context.Context
	Cancel context.CancelFunc
}

func DefaultConfig() *Config {
	ctx, cancel := context.WithCancel(context.WithValue(context.Background(), ctxKeyWaitGroup{}, new(sync.WaitGroup)))
	return &Config{
		Ctx:     ctx,
		Cancel:  cancel,
		AppName: "server",
		DBConfig: dbcore.DBConfig{
			AutoMigrate: true,
		},
		MinioConfig: miniocore.MinioConfig{},
	}
}

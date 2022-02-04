package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/masahide/addsshkey/pkg/opcli"
)

const confPathName = "config"

type Secrets struct {
	Map    map[string]Secret `json:"map"`
	Expiry time.Time         `json:"expiry"`

	AuthWhenAccess bool `json:"auth"`
}

type Secret struct {
	Value  string
	Expiry time.Time
}

type Config struct {
	OpConfig    *opcli.Config
	ConfBaseDir string
}

func New(confBasePath string) (*Config, error) {
	c, err := loadConfig(confBasePath)
	if err != nil {
		c, err = newConfig(confBasePath)
	}
	return c, err
}

func newConfig(confBasePath string) (*Config, error) {
	dir := filepath.Join(confBasePath)
	err := os.MkdirAll(dir, 0700)
	c := &Config{
		OpConfig: opcli.NewConfig(confBasePath),
	}
	if err != nil {
		return c, fmt.Errorf("newConfig err:%w", err)
	}
	return c, err
}

func loadConfig(confBasePath string) (*Config, error) {
	filePath := filepath.Join(confBasePath, confPathName)
	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	c := Config{}
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

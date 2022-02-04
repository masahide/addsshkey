package opcli

import (
	"encoding/json"
	"os"
	"path/filepath"
)

const confPathName = "op_config"

type Config struct {
	CmdPath      string   `json:"cmd_path" default:"op"`
	AccountIDs   []string `json:"accounts"`
	baseDirPath  string
	opConfigPath string
}
type OpAccountInfo struct {
	ID         string `json:"id"`
	URL        string `json:"url"`
	Email      string `json:"email"`
	AccountKey string `json:"account_key"`
}

func NewConfig(baseDirPath string) *Config {
	return &Config{
		CmdPath:      "op",
		AccountIDs:   []string{},
		baseDirPath:  baseDirPath,
		opConfigPath: filepath.Join(baseDirPath, confPathName, "config"),
	}
}

// op cli config file default path: ~/.config/op/config
type opConfigFile struct {
	LatestSignin string      `json:"latest_signin"`
	Device       string      `json:"device"`
	Accounts     []opAccount `json:"accounts"`
}

type opAccount struct {
	Shorthand  string `json:"shorthand"`
	URL        string `json:"url"`
	Email      string `json:"email"`
	AccountKey string `json:"accountKey"`
	UserUUID   string `json:"userUUID"`
	Dsecret    string `json:"dsecret"`
}

func loadOpConfig(confBasePath string) (opConfigFile, error) {
	c := opConfigFile{}
	filePath := filepath.Join(confBasePath, confPathName, "config")
	f, err := os.Open(filePath)
	if err != nil {
		return c, err
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		return c, err
	}
	return c, nil
}

/*
func (c *Config) Write() error {
	f, err := os.OpenFile(c.filePath, os.O_RDWR|os.O_CREATE, 0600)
	if err != nil {
		return err
	}
	defer f.Close()
	return json.NewEncoder(f).Encode(c)
}
*/

package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// BBSConfig stores godis's config
type BBSConfig struct {
	MySQL struct {
		Host     string
		Port     string
		Username string
		Password string
		DbName   string
	}

	OutputToTerminal bool
	LogDir           string

	*viper.Viper
}

var bbsConf *BBSConfig

const (
	confFileDev = "./"
	confFilePro = "/root/work/repos/neighbor_bbs/"
)

// GetConf is used by a function outside the package to get the configuration
func GetConf() *BBSConfig {
	return bbsConf
}

// ParseConf parses config whose format is toml
func ParseConf() {
	v := viper.New()

	v.SetConfigName("bbs")
	v.AddConfigPath(confFileDev)
	v.AddConfigPath(confFilePro)
	// v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config err:%v", err)
		os.Exit(-1)
	}

	v.SetDefault("OutputToTerminal", true)
	v.SetDefault("LogDir", "../log/")

	bbsConf = &BBSConfig{}
	if err := v.Unmarshal(bbsConf); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to unmarshal config :%v", err)
		os.Exit(-1)
	}

	bbsConf.Viper = v
}

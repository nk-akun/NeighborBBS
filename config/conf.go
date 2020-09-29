package config

import (
	"fmt"
	"os"

	"github.com/spf13/viper"
)

// GodisConfig stores godis's config
type GodisConfig struct {
	OutputToTerminal bool
	LogDir           string
	*viper.Viper
}

var godisConf *GodisConfig

const confFile = "/Users/marathon/Work/mygo/src/github.com/nk-akun/NeighborBBS"

// GetConf is used by a function outside the package to get the configuration
func GetConf() *GodisConfig {
	return godisConf
}

// ParseConf parses config whose format is toml
func ParseConf() {
	v := viper.New()

	v.SetConfigName("bbs")
	v.AddConfigPath(confFile)
	// v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load config %s err:%v", confFile, err)
		os.Exit(-1)
	}

	v.SetDefault("OutputToTerminal", true)
	v.SetDefault("LogDir", "../log/")

	godisConf = &GodisConfig{}
	// if err := v.Unmarshal(godisConf); err != nil {
	// 	fmt.Fprintf(os.Stderr, "Failed to unmarshal config :%v", err)
	// 	os.Exit(-1)
	// }

	godisConf.Viper = v
}

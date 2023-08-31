package framework

import (
	"github.com/BurntSushi/toml"
	"log"
)

var gconf *FrameworkConf

type HttpConfig struct {
	Port         int    `toml:"port"`
	StaticDir    string `toml:"staticDir"`
	StaticPrefix string `toml:"staticPrefix"`
}

type HttpsConfig struct {
	Port int    `toml:"port"`
	Cert string `toml:"cert"`
	Key  string `toml:"key"`
}

type XrpcConfig struct {
	ConnectTimeout int    `toml:"connectTimeout"`
	ReadTimeout    int    `toml:"readTimeout"`
	WriteTimeout   int    `toml:"writeTimeout"`
	Server         string `toml:"server"`
}

type FrameworkConf struct {
	Log   LogConfig   `toml:"log"`
	Http  HttpConfig  `toml:"http"`
	Https HttpsConfig `toml:"https"`
	Xrpc  XrpcConfig  `toml:"xrpc"`
}

type LogConfig struct {
	LogDir      string `toml:"logDir"`
	LogFile     string `toml:"logFile"`
	LogLevel    string `toml:"logLevel"`
	LogToStderr bool   `toml:"logToStderr"`
}

func loadConf(confFile string) *FrameworkConf {
	var conf FrameworkConf
	if _, err := toml.DecodeFile(confFile, &conf); err != nil {
		log.Fatalf("Error decoding config file: %s", err)
	}
	gconf = &conf
	return gconf
}

func GetConf() *FrameworkConf {
	return gconf
}

func GetStaticDir() string {
	return gconf.Http.StaticDir
}

func GetXrpcConfig() XrpcConfig {
	return gconf.Xrpc
}

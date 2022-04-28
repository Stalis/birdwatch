package config

import (
	"fmt"
	"os"
	"regexp"
	"time"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/providers/structs"
	"github.com/spf13/pflag"
)

const ErrNoSuchFileOrDirectory = "open .*: no such file or directory"

var config *Config

type Config struct {
	Port   int                 `koanf:"port"`
	Host   string              `koanf:"host"`
	Memory MemoryWatcherConfig `koanf:"memory"`
}

type MemoryWatcherConfig struct {
	Enabled      bool          `koanf:"enabled"`
	ScanInterval time.Duration `koanf:"scan_interval"`
}

func Get() (*Config, error) {
	if config == nil {
		res, err := InitConfig()
		if err != nil {
			return nil, err
		}
		config = res
	}

	return config, nil
}

func InitConfig() (*Config, error) {
	f := initFlagSet()

	k := koanf.New(".")
	k, err := loadDefaultValues(k)
	if err != nil {
		return nil, err
	}

	k, err = initConfiguration(k, f)
	if err != nil {
		return nil, err
	}

	res := &Config{}

	if err := k.Unmarshal("", res); err != nil {
		return nil, err
	}

	return res, nil
}

func initFlagSet() *pflag.FlagSet {
	f := pflag.NewFlagSet("config", pflag.ContinueOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}
	f.IntP("port", "p", 50051, "Port for GRPC server")
	f.StringP("host", "h", "localhost", "Host for GRPC server")
	f.StringP("config", "c", "config.yaml", "config.yaml file path")
	f.Parse(os.Args[1:])
	return f
}

func loadDefaultValues(k *koanf.Koanf) (*koanf.Koanf, error) {
	err := k.Load(structs.Provider(Config{
		Port: 50051,
		Host: "localhost",
		Memory: MemoryWatcherConfig{
			Enabled:      true,
			ScanInterval: time.Second,
		},
	}, ""), nil)
	if err != nil {
		return nil, err
	}

	return k, nil
}

func initConfiguration(k *koanf.Koanf, f *pflag.FlagSet) (*koanf.Koanf, error) {
	confPath, err := f.GetString("config")
	if err != nil {
		return nil, err
	}

	if err := k.Load(file.Provider(confPath), yaml.Parser()); err != nil {
		if match, _ := regexp.MatchString(ErrNoSuchFileOrDirectory, err.Error()); !match {
			return nil, err
		}
	}

	if err := k.Load(posflag.Provider(f, ".", k), nil); err != nil {
		return nil, err
	}

	return k, nil
}

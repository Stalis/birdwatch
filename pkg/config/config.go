package config

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/Stalis/birdwatch/pkg/log"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/file"
	"github.com/knadh/koanf/providers/posflag"
	"github.com/knadh/koanf/providers/structs"
	"github.com/spf13/pflag"
)

const (
	ErrNoSuchFileOrDirectory = "open .*: no such file or directory"
)

const (
	EnvPrefix = "BW_"
)

var config *Config

// Server configuration.
type Config struct {
	// Port with api server is listen
	Port int `koanf:"port" validate:"gt=0,lte=65535"`
	// Host name, using for listening
	Host string `koanf:"host" validate:"required"`
	// Logger configuration
	Logging LogConfig `koanf:"logging" validate:"required"`
	// Memory wather configuration
	Memory MemoryWatcherConfig `koanf:"memory" validate:"required"`
}

// Logger configuration.
type LogConfig struct {
	// Restrict logger print output to console
	Verbose bool `koanf:"verbose"`
	// Logging level
	Level string `koanf:"level" validate:"required"`
	// File to logging output
	File string `koanf:"file" validate:"required"`
}

// Memory watcher configuration.
type MemoryWatcherConfig struct {
	// If false - disabling memory watcher
	Enabled bool `koanf:"enabled"`
	// Interval for scanning memory state
	ScanInterval time.Duration `koanf:"scan_interval"`
}

// Return configuration struct, if initialized,
// else - try to initialize and return configuration struct
// or error if it fails.
func Get() (*Config, error) {
	if config == nil {
		res, err := initConfig()
		if err != nil {
			return nil, err
		}
		config = res
	}

	return config, nil
}

// Try to initialize configuration struct.
func initConfig() (*Config, error) {
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

	if err := validate(res); err != nil {
		fmt.Println(f.FlagUsages())
		return nil, err
	}

	return res, nil
}

func initFlagSet() *pflag.FlagSet {
	f := pflag.NewFlagSet("config", pflag.ExitOnError)
	f.Usage = func() {
		fmt.Println(f.FlagUsages())
		os.Exit(0)
	}

	f.IntP("port", "p", 0, "Port for GRPC server")
	f.StringP("host", "h", "localhost", "Host for GRPC server")
	f.StringP("config", "c", "./config.yaml", "config.yaml file path")
	f.String("logging-level", log.ErrorLevel, fmt.Sprintf("Logging level(one of %v)", log.LevelsList()))
	f.String("logging-file", "server.log", "Logging file")
	f.BoolP("logging-verbose", "v", false, "Verbose mode(logging to stdout/stderr)")

	f.SetNormalizeFunc(wordSeparationNormalizeFunc)

	f.Parse(os.Args[1:])

	return f
}

func loadDefaultValues(k *koanf.Koanf) (*koanf.Koanf, error) {
	err := k.Load(structs.Provider(defaultConfig, "koanf"), nil)
	if err != nil {
		return nil, err
	}

	return k, nil
}

func envVarsProvider() *env.Env {
	return env.Provider(EnvPrefix, ".", func(s string) string {
		return strings.ReplaceAll(strings.ToLower(
			strings.TrimPrefix(s, EnvPrefix)), "_", ".")
	})
}

func initConfiguration(k *koanf.Koanf, f *pflag.FlagSet) (*koanf.Koanf, error) {
	if err := k.Load(envVarsProvider(), nil); err != nil {
		return nil, err
	}

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

func wordSeparationNormalizeFunc(f *pflag.FlagSet, name string) pflag.NormalizedName {
	from := []string{"-", "_"}
	to := "."
	for _, sep := range from {
		name = strings.ReplaceAll(name, sep, to)
	}
	return pflag.NormalizedName(name)
}

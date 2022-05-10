package log

import (
	"os"
	"path/filepath"
)

const (
	ErrorLevel = "Error"
	DebugLevel = "Debug"
	InfoLevel  = "Info"
	WarnLevel  = "Warn"
)

type Config struct {
	Console bool
	Level   string
	File    string
}

func LevelsList() []string {
	return []string{
		ErrorLevel,
		DebugLevel,
		InfoLevel,
		WarnLevel,
	}
}

func openOrCreateFile(path string) (*os.File, error) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModeDir|os.ModePerm); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0o666)
	if err != nil {
		return nil, err
	}

	return file, nil
}

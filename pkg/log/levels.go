package log

const (
	ErrorLevel = "Error"
	DebugLevel = "Debug"
	InfoLevel  = "Info"
	WarnLevel  = "Warn"
)

func LevelsList() []string {
	return []string{
		ErrorLevel,
		DebugLevel,
		InfoLevel,
		WarnLevel,
	}
}

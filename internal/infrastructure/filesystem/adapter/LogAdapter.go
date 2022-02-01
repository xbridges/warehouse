package adapter

type LogAdapter interface {
	Info(message string, fieldmap map[string]interface{})
	Error(message string, fieldmap map[string]interface{})
	Debug(message string, fieldmap map[string]interface{})
	Warn(message string, fieldmap map[string]interface{})
	Panic(message string, fieldmap map[string]interface{})
	Fatal(message string, fieldmap map[string]interface{})
	Sync()
}

package value

type valueLogParameter struct {
	Level 		LogLevel
	Destination LogDestination
	LifeCycle 	LogLifeCycle
}

type LogLevel		string
type LogDestination	string
type LogLifeCycle 	int

type LogParameter interface {
	Params() (string, string, int)
}

func NewLogParameter(level string, path string, life_cycle_day int) LogParameter {
	
	if life_cycle_day < 1 {
		life_cycle_day = 1
	}

	return &valueLogParameter{
		Level: LogLevel(level),
		Destination: LogDestination(path),
		LifeCycle: LogLifeCycle(life_cycle_day),
	}
}

func(v *valueLogParameter) Params() (string, string, int){
	return string(v.Level), string(v.Destination), int(v.LifeCycle)
}